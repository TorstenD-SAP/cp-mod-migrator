//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v294

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditLog) DeepCopyInto(out *AuditLog) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditLog.
func (in *AuditLog) DeepCopy() *AuditLog {
	if in == nil {
		return nil
	}
	out := new(AuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	out.OAuth = in.OAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BusinessDataTunnel) DeepCopyInto(out *BusinessDataTunnel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BusinessDataTunnel.
func (in *BusinessDataTunnel) DeepCopy() *BusinessDataTunnel {
	if in == nil {
		return nil
	}
	out := new(BusinessDataTunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.Integration = in.Integration
	out.ConnectivityService = in.ConnectivityService
	out.MultiRegionMode = in.MultiRegionMode
	out.Servers = in.Servers
	out.ServiceChannels = in.ServiceChannels
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectivityProxy) DeepCopyInto(out *ConnectivityProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectivityProxy.
func (in *ConnectivityProxy) DeepCopy() *ConnectivityProxy {
	if in == nil {
		return nil
	}
	out := new(ConnectivityProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectivityProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectivityProxyList) DeepCopyInto(out *ConnectivityProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConnectivityProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectivityProxyList.
func (in *ConnectivityProxyList) DeepCopy() *ConnectivityProxyList {
	if in == nil {
		return nil
	}
	out := new(ConnectivityProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectivityProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectivityService) DeepCopyInto(out *ConnectivityService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectivityService.
func (in *ConnectivityService) DeepCopy() *ConnectivityService {
	if in == nil {
		return nil
	}
	out := new(ConnectivityService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployment) DeepCopyInto(out *Deployment) {
	*out = *in
	out.RestartWatcher = in.RestartWatcher
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployment.
func (in *Deployment) DeepCopy() *Deployment {
	if in == nil {
		return nil
	}
	out := new(Deployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.Selector = in.Selector
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTP) DeepCopyInto(out *HTTP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP.
func (in *HTTP) DeepCopy() *HTTP {
	if in == nil {
		return nil
	}
	out := new(HTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	out.Tls = in.Tls
	out.Timeouts = in.Timeouts
	in.Istio.DeepCopyInto(&out.Istio)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTls) DeepCopyInto(out *IngressTls) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTls.
func (in *IngressTls) DeepCopy() *IngressTls {
	if in == nil {
		return nil
	}
	out := new(IngressTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Integration) DeepCopyInto(out *Integration) {
	*out = *in
	out.AuditLog = in.AuditLog
	out.ConnectivityService = in.ConnectivityService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Integration.
func (in *Integration) DeepCopy() *Integration {
	if in == nil {
		return nil
	}
	out := new(Integration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Istio) DeepCopyInto(out *Istio) {
	*out = *in
	out.Gateway = in.Gateway
	in.Tls.DeepCopyInto(&out.Tls)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Istio.
func (in *Istio) DeepCopy() *Istio {
	if in == nil {
		return nil
	}
	out := new(Istio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioTls) DeepCopyInto(out *IstioTls) {
	*out = *in
	if in.Ciphers != nil {
		in, out := &in.Ciphers, &out.Ciphers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioTls.
func (in *IstioTls) DeepCopy() *IstioTls {
	if in == nil {
		return nil
	}
	out := new(IstioTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRegionMode) DeepCopyInto(out *MultiRegionMode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRegionMode.
func (in *MultiRegionMode) DeepCopy() *MultiRegionMode {
	if in == nil {
		return nil
	}
	out := new(MultiRegionMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth) DeepCopyInto(out *OAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth.
func (in *OAuth) DeepCopy() *OAuth {
	if in == nil {
		return nil
	}
	out := new(OAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
	out.Authorization = in.Authorization
	out.HTTP = in.HTTP
	out.RfcAndLdap = in.RfcAndLdap
	out.Socks5 = in.Socks5
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyCfg) DeepCopyInto(out *ProxyCfg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyCfg.
func (in *ProxyCfg) DeepCopy() *ProxyCfg {
	if in == nil {
		return nil
	}
	out := new(ProxyCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestartWatcher) DeepCopyInto(out *RestartWatcher) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestartWatcher.
func (in *RestartWatcher) DeepCopy() *RestartWatcher {
	if in == nil {
		return nil
	}
	out := new(RestartWatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RfcAndLdap) DeepCopyInto(out *RfcAndLdap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RfcAndLdap.
func (in *RfcAndLdap) DeepCopy() *RfcAndLdap {
	if in == nil {
		return nil
	}
	out := new(RfcAndLdap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretConfig) DeepCopyInto(out *SecretConfig) {
	*out = *in
	out.Integration = in.Integration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretConfig.
func (in *SecretConfig) DeepCopy() *SecretConfig {
	if in == nil {
		return nil
	}
	out := new(SecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretConfigIntegration) DeepCopyInto(out *SecretConfigIntegration) {
	*out = *in
	out.ConnectivityService = in.ConnectivityService
	out.AuditLogService = in.AuditLogService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretConfigIntegration.
func (in *SecretConfigIntegration) DeepCopy() *SecretConfigIntegration {
	if in == nil {
		return nil
	}
	out := new(SecretConfigIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers) DeepCopyInto(out *Servers) {
	*out = *in
	out.BusinessDataTunnel = in.BusinessDataTunnel
	out.Proxy = in.Proxy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers.
func (in *Servers) DeepCopy() *Servers {
	if in == nil {
		return nil
	}
	out := new(Servers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceChannels) DeepCopyInto(out *ServiceChannels) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceChannels.
func (in *ServiceChannels) DeepCopy() *ServiceChannels {
	if in == nil {
		return nil
	}
	out := new(ServiceChannels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSecretConfig) DeepCopyInto(out *ServiceSecretConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSecretConfig.
func (in *ServiceSecretConfig) DeepCopy() *ServiceSecretConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Socks5) DeepCopyInto(out *Socks5) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Socks5.
func (in *Socks5) DeepCopy() *Socks5 {
	if in == nil {
		return nil
	}
	out := new(Socks5)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	out.Config = in.Config
	out.Deployment = in.Deployment
	in.Ingress.DeepCopyInto(&out.Ingress)
	out.SecretConfig = in.SecretConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutProxy) DeepCopyInto(out *TimeoutProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutProxy.
func (in *TimeoutProxy) DeepCopy() *TimeoutProxy {
	if in == nil {
		return nil
	}
	out := new(TimeoutProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
	out.Proxy = in.Proxy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}
