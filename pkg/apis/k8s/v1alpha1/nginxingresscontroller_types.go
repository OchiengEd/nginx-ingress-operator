package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NginxIngressControllerSpec defines the desired state of NginxIngressController
// +operator-sdk:gen-csv:customresourcedefinitions.resources=`Deployment,v1,"nginx-ingress-operator"`
type NginxIngressControllerSpec struct {
	// The type of the Ingress Controller installation - deployment or daemonset.
	// +kubebuilder:validation:Enum=deployment;daemonset
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	Type string `json:"type"`
	// Deploys the Ingress Controller for NGINX Plus. The default is false meaning the Ingress Controller will be deployed for NGINX OSS.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	NginxPlus bool `json:"nginxPlus"`
	// The image of the Ingress Controller.
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	Image Image `json:"image"`
	// The number of replicas of the Ingress Controller pod. The default is 1. Only applies if the type is set to deployment.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	Replicas *int32 `json:"replicas"`
	// The TLS Secret for TLS termination of the default server. The format is namespace/name.
	// If not specified, the operator will generate and deploy a TLS Secret with a self-signed certificate and key.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	DefaultSecret string `json:"defaultSecret"`
	// The type of the Service for the Ingress Controller. Valid Service types are: NodePort and LoadBalancer.
	// +kubebuilder:validation:Enum=NodePort;LoadBalancer
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	ServiceType string `json:"serviceType"`
	// Enables the use of NGINX Ingress Resource Definitions (VirtualServer and VirtualServerRoute).
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	EnableCRDs bool `json:"enableCRDs"`
	// Enable custom NGINX configuration snippets in VirtualServer and VirtualServerRoute resources.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	EnableSnippets bool `json:"enableSnippets"`
	// +kubebuilder:validation:Optional
	// A class of the Ingress controller. The Ingress controller only processes Ingress resources that belong to its
	// class (in other words, have the annotation “kubernetes.io/ingress.class”).
	// Additionally, the Ingress controller processes Ingress resources that do not have that annotation,
	// which can be disabled by setting UseIngressClassOnly to true. Default is `nginx`.
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	IngressClass string `json:"ingressClass"`
	// Ignore Ingress resources without the “kubernetes.io/ingress.class” annotation.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	UseIngressClassOnly bool `json:"useIngressClassOnly"`
	// Namespace to watch for Ingress resources. By default the Ingress controller watches all namespaces.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	WatchNamespace string `json:"watchNamespace"`
	// Adds a new location to the default server. The location responds with the 200 status code for any request.
	// Useful for external health-checking of the Ingress controller.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	HealthStatus *HealthStatus `json:"healthStatus,omitempty"`
	// Enable debugging for NGINX. Uses the nginx-debug binary. Requires ‘error-log-level: debug’ in the ConfigMapData.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	NginxDebug bool `json:"nginxDebug"`
	// Log level for V logs.
	// Format is 0 - 3
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	LogLevel uint8 `json:"logLevel"`
	// NGINX stub_status, or the NGINX Plus API.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	NginxStatus *NginxStatus `json:"nginxStatus,omitempty"`
	// Update the address field in the status of Ingresses resources.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	ReportIngressStatus *ReportIngressStatus `json:"reportIngressStatus,omitempty"`
	// Enables Leader election to avoid multiple replicas of the controller reporting the status of Ingress resources
	// – only one replica will report status.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	EnableLeaderElection bool `json:"enableLeaderElection"`
	// A Secret with a TLS certificate and key for TLS termination of every Ingress host for which TLS termination is enabled but the Secret is not specified.
	// If the argument is not set, for such Ingress hosts NGINX will break any attempt to establish a TLS connection.
	// If the argument is set, but the Ingress controller is not able to fetch the Secret from Kubernetes API, the Ingress Controller will fail to start.
	// Format is namespace/name.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	WildcardTLS string `json:"wildcardTLS"`
	// NGINX or NGINX Plus metrics in the Prometheus format.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	Prometheus *Prometheus `json:"prometheus,omitempty"`
	// Initial values of the Ingress Controller ConfigMap.
	// Check https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/ for
	// more information about possible values.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	ConfigMapData map[string]string `json:"configMapData,omitempty"`
	// The GlobalConfiguration resource for global configuration of the Ingress Controller.
	// Format is namespace/name.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	GlobalConfiguration string `json:"globalConfiguration"`
	// Enable TLS Passthrough on port 443.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	EnableTLSPassthrough bool `json:"enableTLSPassthrough"`
	// App Protect support configuration.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	AppProtect *AppProtect `json:"appProtect"`
	// Timeout in milliseconds which the Ingress Controller will wait for a successful NGINX reload after a change or at the initial start.
	// +kubebuilder:validation:Optional
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	NginxReloadTimeout int `json:"nginxReloadTimeout"`
}

// Image defines the Repository, Tag and ImagePullPolicy of the Ingress Controller Image.
type Image struct {
	// The repository of the image.
	Repository string `json:"repository"`
	// The tag (version) of the image.
	Tag string `json:"tag"`
	// The ImagePullPolicy of the image.
	// +kubebuilder:validation:Enum=Never;Always;IfNotPresent
	PullPolicy string `json:"pullPolicy"`
}

// HealthStatus defines the health status of the Ingress Controller.
type HealthStatus struct {
	// Enable the HealthStatus.
	Enable bool `json:"enable"`
	// URI of the location. Default is `/nginx-health`.
	// +kubebuilder:validation:Optional
	URI string `json:"uri"`
}

// NginxStatus defines the NGINX Status of the Ingress Controller.
type NginxStatus struct {
	// Enable the NginxStatus.
	Enable bool `json:"enable"`
	// Set the port where the NGINX stub_status or the NGINX Plus API is exposed. Default is 8080.
	// Format is 1023 - 65535
	// +kubebuilder:validation:Minimum=1023
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	// +nullable
	Port *uint16 `json:"port"`
	// Whitelist IPv4 IP/CIDR blocks to allow access to NGINX stub_status or the NGINX Plus API.
	// Separate multiple IP/CIDR by commas. (default “127.0.0.1”)
	// +kubebuilder:validation:Optional
	AllowCidrs string `json:"allowCidrs"`
}

// ReportIngressStatus defines the report of the status of the Ingress Resources.
type ReportIngressStatus struct {
	// Enable the ReportIngressStatus.
	Enable bool `json:"enable"`
	// Specifies the name of the service with the type LoadBalancer through which the Ingress controller pods are exposed externally.
	// The external address of the service is used when reporting the status of Ingress resources.
	// Note: Only if ServiceType is different than LoadBalancer.
	// +kubebuilder:validation:Optional
	ExternalService string `json:"externalService"`
}

// Prometheus defines the Prometheus metrics for the Ingress Controller.
type Prometheus struct {
	// Enable Prometheus metrics.
	Enable bool `json:"enable"`
	// Sets the port where the Prometheus metrics are exposed. Default is 9113.
	// Format is 1023 - 65535
	// +kubebuilder:validation:Minimum=1023
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	// +nullable
	Port *uint16 `json:"port"`
}

// App Protect support configuration.
type AppProtect struct {
	// Enable App Protect.
	Enable bool `json:"enable"`
}

// NginxIngressControllerStatus defines the observed state of NginxIngressController
type NginxIngressControllerStatus struct {
	// Deployed is true if the Operator has finished the deployment of the NginxIngressController.
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.statusDescriptors=true
	Deployed bool `json:"deployed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NginxIngressController is the Schema for the nginxingresscontrollers API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nginxingresscontrollers,scope=Namespaced
type NginxIngressController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NginxIngressControllerSpec   `json:"spec,omitempty"`
	Status NginxIngressControllerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NginxIngressControllerList contains a list of NginxIngressController
type NginxIngressControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NginxIngressController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NginxIngressController{}, &NginxIngressControllerList{})
}
