// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerWebhookCustomizationSpec) DeepCopyInto(out *AdmissionControllerWebhookCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerWebhookCustomizationSpec.
func (in *AdmissionControllerWebhookCustomizationSpec) DeepCopy() *AdmissionControllerWebhookCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerWebhookCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerWebhookStatus) DeepCopyInto(out *AdmissionControllerWebhookStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerWebhookStatus.
func (in *AdmissionControllerWebhookStatus) DeepCopy() *AdmissionControllerWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerWebhookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsodyStatus) DeepCopyInto(out *AppsodyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsodyStatus.
func (in *AppsodyStatus) DeepCopy() *AppsodyStatus {
	if in == nil {
		return nil
	}
	out := new(AppsodyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheCustomizationSpec) DeepCopyInto(out *CheCustomizationSpec) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	out.CheOperatorInstance = in.CheOperatorInstance
	out.KabaneroChe = in.KabaneroChe
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheCustomizationSpec.
func (in *CheCustomizationSpec) DeepCopy() *CheCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(CheCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheOperatorInstanceSpec) DeepCopyInto(out *CheOperatorInstanceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheOperatorInstanceSpec.
func (in *CheOperatorInstanceSpec) DeepCopy() *CheOperatorInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CheOperatorInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheOperatorStatus) DeepCopyInto(out *CheOperatorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheOperatorStatus.
func (in *CheOperatorStatus) DeepCopy() *CheOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(CheOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheStatus) DeepCopyInto(out *CheStatus) {
	*out = *in
	out.CheOperator = in.CheOperator
	out.KabaneroChe = in.KabaneroChe
	out.KabaneroCheInstance = in.KabaneroCheInstance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheStatus.
func (in *CheStatus) DeepCopy() *CheStatus {
	if in == nil {
		return nil
	}
	out := new(CheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CliStatus) DeepCopyInto(out *CliStatus) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CliStatus.
func (in *CliStatus) DeepCopy() *CliStatus {
	if in == nil {
		return nil
	}
	out := new(CliStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collection) DeepCopyInto(out *Collection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collection.
func (in *Collection) DeepCopy() *Collection {
	if in == nil {
		return nil
	}
	out := new(Collection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionList) DeepCopyInto(out *CollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionList.
func (in *CollectionList) DeepCopy() *CollectionList {
	if in == nil {
		return nil
	}
	out := new(CollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionSpec) DeepCopyInto(out *CollectionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionSpec.
func (in *CollectionSpec) DeepCopy() *CollectionSpec {
	if in == nil {
		return nil
	}
	out := new(CollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionStatus) DeepCopyInto(out *CollectionStatus) {
	*out = *in
	if in.ActivePipelines != nil {
		in, out := &in.ActivePipelines, &out.ActivePipelines
		*out = make([]PipelineStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]Image, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionStatus.
func (in *CollectionStatus) DeepCopy() *CollectionStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubConfig) DeepCopyInto(out *GithubConfig) {
	*out = *in
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubConfig.
func (in *GithubConfig) DeepCopy() *GithubConfig {
	if in == nil {
		return nil
	}
	out := new(GithubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCollectionConfig) DeepCopyInto(out *InstanceCollectionConfig) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]RepositoryConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCollectionConfig.
func (in *InstanceCollectionConfig) DeepCopy() *InstanceCollectionConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceCollectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kabanero) DeepCopyInto(out *Kabanero) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kabanero.
func (in *Kabanero) DeepCopy() *Kabanero {
	if in == nil {
		return nil
	}
	out := new(Kabanero)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kabanero) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroCheInstanceStatus) DeepCopyInto(out *KabaneroCheInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroCheInstanceStatus.
func (in *KabaneroCheInstanceStatus) DeepCopy() *KabaneroCheInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroCheInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroCheSpec) DeepCopyInto(out *KabaneroCheSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroCheSpec.
func (in *KabaneroCheSpec) DeepCopy() *KabaneroCheSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroCheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroCheStatus) DeepCopyInto(out *KabaneroCheStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroCheStatus.
func (in *KabaneroCheStatus) DeepCopy() *KabaneroCheStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroCheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroCliServicesCustomizationSpec) DeepCopyInto(out *KabaneroCliServicesCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroCliServicesCustomizationSpec.
func (in *KabaneroCliServicesCustomizationSpec) DeepCopy() *KabaneroCliServicesCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroCliServicesCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroInstanceStatus) DeepCopyInto(out *KabaneroInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroInstanceStatus.
func (in *KabaneroInstanceStatus) DeepCopy() *KabaneroInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroLandingCustomizationSpec) DeepCopyInto(out *KabaneroLandingCustomizationSpec) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroLandingCustomizationSpec.
func (in *KabaneroLandingCustomizationSpec) DeepCopy() *KabaneroLandingCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroLandingCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroLandingPageStatus) DeepCopyInto(out *KabaneroLandingPageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroLandingPageStatus.
func (in *KabaneroLandingPageStatus) DeepCopy() *KabaneroLandingPageStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroLandingPageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroList) DeepCopyInto(out *KabaneroList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kabanero, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroList.
func (in *KabaneroList) DeepCopy() *KabaneroList {
	if in == nil {
		return nil
	}
	out := new(KabaneroList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KabaneroList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroSpec) DeepCopyInto(out *KabaneroSpec) {
	*out = *in
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Github.DeepCopyInto(&out.Github)
	in.Collections.DeepCopyInto(&out.Collections)
	out.Tekton = in.Tekton
	out.CliServices = in.CliServices
	in.Landing.DeepCopyInto(&out.Landing)
	in.Che.DeepCopyInto(&out.Che)
	out.Webhook = in.Webhook
	out.AdmissionControllerWebhook = in.AdmissionControllerWebhook
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroSpec.
func (in *KabaneroSpec) DeepCopy() *KabaneroSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroStatus) DeepCopyInto(out *KabaneroStatus) {
	*out = *in
	out.KabaneroInstance = in.KabaneroInstance
	out.KnativeEventing = in.KnativeEventing
	out.Serverless = in.Serverless
	out.Tekton = in.Tekton
	in.Cli.DeepCopyInto(&out.Cli)
	if in.Landing != nil {
		in, out := &in.Landing, &out.Landing
		*out = new(KabaneroLandingPageStatus)
		**out = **in
	}
	out.Appsody = in.Appsody
	if in.Kappnav != nil {
		in, out := &in.Kappnav, &out.Kappnav
		*out = new(KappnavStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Che != nil {
		in, out := &in.Che, &out.Che
		*out = new(CheStatus)
		**out = **in
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookStatus)
		(*in).DeepCopyInto(*out)
	}
	out.AdmissionControllerWebhook = in.AdmissionControllerWebhook
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroStatus.
func (in *KabaneroStatus) DeepCopy() *KabaneroStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappnavStatus) DeepCopyInto(out *KappnavStatus) {
	*out = *in
	if in.UiLocations != nil {
		in, out := &in.UiLocations, &out.UiLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ApiLocations != nil {
		in, out := &in.ApiLocations, &out.ApiLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappnavStatus.
func (in *KappnavStatus) DeepCopy() *KappnavStatus {
	if in == nil {
		return nil
	}
	out := new(KappnavStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventingStatus) DeepCopyInto(out *KnativeEventingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventingStatus.
func (in *KnativeEventingStatus) DeepCopy() *KnativeEventingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeEventingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingStatus) DeepCopyInto(out *KnativeServingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingStatus.
func (in *KnativeServingStatus) DeepCopy() *KnativeServingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeServingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	if in.ActiveAssets != nil {
		in, out := &in.ActiveAssets, &out.ActiveAssets
		*out = make([]RepositoryAssetStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryAssetStatus) DeepCopyInto(out *RepositoryAssetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryAssetStatus.
func (in *RepositoryAssetStatus) DeepCopy() *RepositoryAssetStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryConfig) DeepCopyInto(out *RepositoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryConfig.
func (in *RepositoryConfig) DeepCopy() *RepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(RepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessStatus) DeepCopyInto(out *ServerlessStatus) {
	*out = *in
	out.KnativeServing = in.KnativeServing
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessStatus.
func (in *ServerlessStatus) DeepCopy() *ServerlessStatus {
	if in == nil {
		return nil
	}
	out := new(ServerlessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonCustomizationSpec) DeepCopyInto(out *TektonCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonCustomizationSpec.
func (in *TektonCustomizationSpec) DeepCopy() *TektonCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(TektonCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonStatus) DeepCopyInto(out *TektonStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonStatus.
func (in *TektonStatus) DeepCopy() *TektonStatus {
	if in == nil {
		return nil
	}
	out := new(TektonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookCustomizationSpec) DeepCopyInto(out *WebhookCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookCustomizationSpec.
func (in *WebhookCustomizationSpec) DeepCopy() *WebhookCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookStatus) DeepCopyInto(out *WebhookStatus) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookStatus.
func (in *WebhookStatus) DeepCopy() *WebhookStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookStatus)
	in.DeepCopyInto(out)
	return out
}
