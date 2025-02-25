package v1

import (
	"fmt"
	"math"
	"strconv"

	ctypes "github.com/ovrclk/akash/provider/cluster/types"
	clusterutil "github.com/ovrclk/akash/provider/cluster/util"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ovrclk/akash/manifest"
	types "github.com/ovrclk/akash/types/v1beta2"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Manifest store metadata, specifications and status of the Lease
type Manifest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ManifestSpec   `json:"spec,omitempty"`
	Status ManifestStatus `json:"status,omitempty"`
}

// ManifestStatus stores state and message of manifest
type ManifestStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// ManifestSpec stores LeaseID, Group and metadata details
type ManifestSpec struct {
	LeaseID LeaseID       `json:"lease_id"`
	Group   ManifestGroup `json:"group"`
}

// Deployment returns the cluster.Deployment that the saved manifest represents.
func (m Manifest) Deployment() (ctypes.Deployment, error) {
	lid, err := m.Spec.LeaseID.toAkash()
	if err != nil {
		return nil, err
	}

	group, err := m.Spec.Group.toAkash()
	if err != nil {
		return nil, err
	}
	return deployment{lid: lid, group: group}, nil
}

type deployment struct {
	lid   mtypes.LeaseID
	group manifest.Group
}

func (d deployment) LeaseID() mtypes.LeaseID {
	return d.lid
}

func (d deployment) ManifestGroup() manifest.Group {
	return d.group
}

// NewManifest creates new manifest with provided details. Returns error in case of failure.
func NewManifest(ns string, lid mtypes.LeaseID, mgroup *manifest.Group) (*Manifest, error) {
	group, err := manifestGroupFromAkash(mgroup)
	if err != nil {
		return nil, err
	}

	return &Manifest{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Manifest",
			APIVersion: "akash.network/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterutil.LeaseIDToNamespace(lid),
			Namespace: ns,
		},
		Spec: ManifestSpec{
			Group:   group,
			LeaseID: leaseIDFromAkash(lid),
		},
	}, nil
}

// LeaseID stores deployment, group sequence, order, provider and metadata
type LeaseID struct {
	Owner    string `json:"owner"`
	DSeq     string `json:"dseq"`
	GSeq     uint32 `json:"gseq"`
	OSeq     uint32 `json:"oseq"`
	Provider string `json:"provider"`
}

// ToAkash returns LeaseID from LeaseID details
func (id LeaseID) toAkash() (mtypes.LeaseID, error) {
	owner, err := sdk.AccAddressFromBech32(id.Owner)
	if err != nil {
		return mtypes.LeaseID{}, err
	}

	provider, err := sdk.AccAddressFromBech32(id.Provider)
	if err != nil {
		return mtypes.LeaseID{}, err
	}

	dseq, err := strconv.ParseUint(id.DSeq, 10, 64)
	if err != nil {
		return mtypes.LeaseID{}, err
	}

	return mtypes.LeaseID{
		Owner:    owner.String(),
		DSeq:     dseq,
		GSeq:     id.GSeq,
		OSeq:     id.OSeq,
		Provider: provider.String(),
	}, nil
}

// LeaseIDFromAkash returns LeaseID instance from akash
func leaseIDFromAkash(id mtypes.LeaseID) LeaseID {
	return LeaseID{
		Owner:    id.Owner,
		DSeq:     strconv.FormatUint(id.DSeq, 10),
		GSeq:     id.GSeq,
		OSeq:     id.OSeq,
		Provider: id.Provider,
	}
}

// ManifestGroup stores metadata, name and list of SDL manifest services
type ManifestGroup struct {
	// Placement profile name
	Name string `json:"name,omitempty"`
	// Service definitions
	Services []ManifestService `json:"services,omitempty"`
}

// ToAkash returns akash group details formatted from manifest group
func (m ManifestGroup) toAkash() (manifest.Group, error) {
	am := manifest.Group{
		Name:     m.Name,
		Services: make([]manifest.Service, 0, len(m.Services)),
	}

	for _, svc := range m.Services {
		asvc, err := svc.toAkash()
		if err != nil {
			return am, err
		}
		am.Services = append(am.Services, asvc)
	}

	return am, nil
}

// ManifestGroupFromAkash returns manifest group instance from akash group
func manifestGroupFromAkash(m *manifest.Group) (ManifestGroup, error) {
	ma := ManifestGroup{
		Name:     m.Name,
		Services: make([]ManifestService, 0, len(m.Services)),
	}

	for _, svc := range m.Services {
		service, err := manifestServiceFromAkash(svc)
		if err != nil {
			return ManifestGroup{}, err
		}

		ma.Services = append(ma.Services, service)
	}

	return ma, nil
}

type ManifestStorageParams struct {
	Name     string `json:"name" yaml:"name"`
	Mount    string `json:"mount" yaml:"mount"`
	ReadOnly bool   `json:"readOnly" yaml:"readOnly"`
}

type ManifestServiceParams struct {
	Storage []ManifestStorageParams `json:"storage,omitempty"`
}

// ManifestService stores name, image, args, env, unit, count and expose list of service
type ManifestService struct {
	// Service name
	Name string `json:"name,omitempty"`
	// Docker image
	Image string   `json:"image,omitempty"`
	Args  []string `json:"args,omitempty"`
	Env   []string `json:"env,omitempty"`
	// Resource requirements
	// in current version of CRD it is named as unit
	Resources ResourceUnits `json:"unit"`
	// Number of instances
	Count uint32 `json:"count,omitempty"`
	// Overlay Network Links
	Expose []ManifestServiceExpose `json:"expose,omitempty"`
	// Miscellaneous service parameters
	Params *ManifestServiceParams `json:"params,omitempty"`
}

func (ms ManifestService) toAkash() (manifest.Service, error) {
	res, err := ms.Resources.toAkash()
	if err != nil {
		return manifest.Service{}, err
	}

	ams := &manifest.Service{
		Name:      ms.Name,
		Image:     ms.Image,
		Args:      ms.Args,
		Env:       ms.Env,
		Resources: res,
		Count:     ms.Count,
		Expose:    make([]manifest.ServiceExpose, 0, len(ms.Expose)),
	}

	for _, expose := range ms.Expose {
		value, err := expose.toAkash()
		if err != nil {
			return manifest.Service{}, err
		}
		ams.Expose = append(ams.Expose, value)
	}

	if ms.Params != nil {
		ams.Params = &manifest.ServiceParams{
			Storage: make([]manifest.StorageParams, 0, len(ms.Params.Storage)),
		}

		for _, storage := range ms.Params.Storage {
			ams.Params.Storage = append(ams.Params.Storage, manifest.StorageParams{
				Name:     storage.Name,
				Mount:    storage.Mount,
				ReadOnly: storage.ReadOnly,
			})
		}
	}

	return *ams, nil
}

func manifestServiceFromAkash(ams manifest.Service) (ManifestService, error) {
	resources, err := resourceUnitsFromAkash(ams.Resources)
	if err != nil {
		return ManifestService{}, err
	}

	ms := ManifestService{
		Name:      ams.Name,
		Image:     ams.Image,
		Args:      ams.Args,
		Env:       ams.Env,
		Resources: resources,
		Count:     ams.Count,
		Expose:    make([]ManifestServiceExpose, 0, len(ams.Expose)),
	}

	for _, expose := range ams.Expose {
		ms.Expose = append(ms.Expose, manifestServiceExposeFromAkash(expose))
	}

	if ams.Params != nil {
		ms.Params = &ManifestServiceParams{
			Storage: make([]ManifestStorageParams, 0, len(ams.Params.Storage)),
		}

		for _, storage := range ams.Params.Storage {
			ms.Params.Storage = append(ms.Params.Storage, ManifestStorageParams{
				Name:     storage.Name,
				Mount:    storage.Mount,
				ReadOnly: storage.ReadOnly,
			})
		}
	}

	return ms, nil
}

// ManifestServiceExpose stores exposed ports and accepted hosts details
type ManifestServiceExpose struct {
	Port         uint16 `json:"port,omitempty"`
	ExternalPort uint16 `json:"external_port,omitempty"`
	Proto        string `json:"proto,omitempty"`
	Service      string `json:"service,omitempty"`
	Global       bool   `json:"global,omitempty"`
	// accepted hostnames
	Hosts       []string                         `json:"hosts,omitempty"`
	HTTPOptions ManifestServiceExposeHTTPOptions `json:"http_options,omitempty"`
}

type ManifestServiceExposeHTTPOptions struct {
	MaxBodySize uint32   `json:"max_body_size,omitempty"`
	ReadTimeout uint32   `json:"read_timeout,omitempty"`
	SendTimeout uint32   `json:"send_timeout,omitempty"`
	NextTries   uint32   `json:"next_tries,omitempty"`
	NextTimeout uint32   `json:"next_timeout,omitempty"`
	NextCases   []string `json:"next_cases,omitempty"`
}

func (mse ManifestServiceExpose) toAkash() (manifest.ServiceExpose, error) {
	proto, err := manifest.ParseServiceProtocol(mse.Proto)
	if err != nil {
		fmt.Printf("foobar: %q\n", mse.Proto)
		return manifest.ServiceExpose{}, err
	}
	return manifest.ServiceExpose{
		Port:         mse.Port,
		ExternalPort: mse.ExternalPort,
		Proto:        proto,
		Service:      mse.Service,
		Global:       mse.Global,
		Hosts:        mse.Hosts,
	}, nil
}

func manifestServiceExposeFromAkash(amse manifest.ServiceExpose) ManifestServiceExpose {
	return ManifestServiceExpose{
		Port:         amse.Port,
		ExternalPort: amse.ExternalPort,
		Proto:        amse.Proto.ToString(),
		Service:      amse.Service,
		Global:       amse.Global,
		Hosts:        amse.Hosts,
		HTTPOptions: ManifestServiceExposeHTTPOptions{
			MaxBodySize: amse.HTTPOptions.MaxBodySize,
			ReadTimeout: amse.HTTPOptions.ReadTimeout,
			SendTimeout: amse.HTTPOptions.SendTimeout,
			NextTries:   amse.HTTPOptions.NextTries,
			NextTimeout: amse.HTTPOptions.NextTimeout,
			NextCases:   amse.HTTPOptions.NextCases,
		},
	}
}

// ResourceUnits stores cpu, memory and storage details
type ResourceUnits struct {
	CPU     uint32   `json:"cpu,omitempty"`
	Memory  string   `json:"memory,omitempty"`
	Storage []string `json:"storage,omitempty"`
}

func (ru ResourceUnits) toAkash() (types.ResourceUnits, error) {
	memory, err := strconv.ParseUint(ru.Memory, 10, 64)
	if err != nil {
		return types.ResourceUnits{}, err
	}

	storage := make([]types.Storage, 0, len(ru.Storage))

	for _, st := range ru.Storage {
		size, err := strconv.ParseUint(st, 10, 64)
		if err != nil {
			return types.ResourceUnits{}, err
		}

		storage = append(storage, types.Storage{
			Quantity: types.NewResourceValue(size),
		})
	}

	return types.ResourceUnits{
		CPU: &types.CPU{
			Units: types.NewResourceValue(uint64(ru.CPU)),
		},
		Memory: &types.Memory{
			Quantity: types.NewResourceValue(memory),
		},
		Storage: storage,
	}, nil
}

func resourceUnitsFromAkash(aru types.ResourceUnits) (ResourceUnits, error) {
	res := ResourceUnits{}
	if aru.CPU != nil {
		if aru.CPU.Units.Value() > math.MaxUint32 {
			return ResourceUnits{}, errors.Errorf("k8s api: cpu units value overflows uint32")
		}
		res.CPU = uint32(aru.CPU.Units.Value())
	}
	if aru.Memory != nil {
		res.Memory = strconv.FormatUint(aru.Memory.Quantity.Value(), 10)
	}

	res.Storage = make([]string, 0, len(aru.Storage))
	for _, size := range aru.Storage {
		res.Storage = append(res.Storage, strconv.FormatUint(size.Quantity.Value(), 10))
	}

	return res, nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManifestList stores metadata and items list of manifest
type ManifestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Manifest `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProviderHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ProviderHostSpec   `json:"spec,omitempty"`
	Status ProviderHostStatus `json:"status,omitempty"`
}

type ProviderHostStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

type ProviderHostSpec struct {
	Owner        string `json:"owner"`
	Provider     string `json:"provider"`
	Hostname     string `json:"hostname"`
	Dseq         uint64 `json:"dseq"`
	Gseq         uint32 `json:"gseq"`
	Oseq         uint32 `json:"oseq"`
	ServiceName  string `json:"service_name"`
	ExternalPort uint32 `json:"external_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProviderHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ProviderHost `json:"items"`
}
