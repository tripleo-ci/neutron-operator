// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronOvsAgent) DeepCopyInto(out *NeutronOvsAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronOvsAgent.
func (in *NeutronOvsAgent) DeepCopy() *NeutronOvsAgent {
	if in == nil {
		return nil
	}
	out := new(NeutronOvsAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeutronOvsAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronOvsAgentList) DeepCopyInto(out *NeutronOvsAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NeutronOvsAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronOvsAgentList.
func (in *NeutronOvsAgentList) DeepCopy() *NeutronOvsAgentList {
	if in == nil {
		return nil
	}
	out := new(NeutronOvsAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeutronOvsAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronOvsAgentSpec) DeepCopyInto(out *NeutronOvsAgentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronOvsAgentSpec.
func (in *NeutronOvsAgentSpec) DeepCopy() *NeutronOvsAgentSpec {
	if in == nil {
		return nil
	}
	out := new(NeutronOvsAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronOvsAgentStatus) DeepCopyInto(out *NeutronOvsAgentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronOvsAgentStatus.
func (in *NeutronOvsAgentStatus) DeepCopy() *NeutronOvsAgentStatus {
	if in == nil {
		return nil
	}
	out := new(NeutronOvsAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSriovAgent) DeepCopyInto(out *NeutronSriovAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSriovAgent.
func (in *NeutronSriovAgent) DeepCopy() *NeutronSriovAgent {
	if in == nil {
		return nil
	}
	out := new(NeutronSriovAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeutronSriovAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSriovAgentList) DeepCopyInto(out *NeutronSriovAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NeutronSriovAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSriovAgentList.
func (in *NeutronSriovAgentList) DeepCopy() *NeutronSriovAgentList {
	if in == nil {
		return nil
	}
	out := new(NeutronSriovAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeutronSriovAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSriovAgentSpec) DeepCopyInto(out *NeutronSriovAgentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSriovAgentSpec.
func (in *NeutronSriovAgentSpec) DeepCopy() *NeutronSriovAgentSpec {
	if in == nil {
		return nil
	}
	out := new(NeutronSriovAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSriovAgentStatus) DeepCopyInto(out *NeutronSriovAgentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSriovAgentStatus.
func (in *NeutronSriovAgentStatus) DeepCopy() *NeutronSriovAgentStatus {
	if in == nil {
		return nil
	}
	out := new(NeutronSriovAgentStatus)
	in.DeepCopyInto(out)
	return out
}
