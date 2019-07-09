package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	auth_v1 "k8s.io/api/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/kubernetes"
	"github.com/kiali/kiali/kubernetes/kubetest"
	"github.com/kiali/kiali/models"
	"github.com/kiali/kiali/tests/data"
)

func TestGetIstioConfigList(t *testing.T) {
	assert := assert.New(t)
	criteria := IstioConfigCriteria{
		Namespace:               "test",
		IncludeGateways:         false,
		IncludeVirtualServices:  false,
		IncludeDestinationRules: false,
		IncludeServiceEntries:   false,
		IncludeRules:            false,
		IncludeAdapters:         false,
		IncludeTemplates:        false,
		IncludeQuotaSpecs:       false,
	}

	configService := mockGetIstioConfigList()

	istioconfigList, err := configService.GetIstioConfigList(criteria)

	assert.Equal(0, len(istioconfigList.Gateways))
	assert.Equal(0, len(istioconfigList.VirtualServices.Items))
	assert.Equal(0, len(istioconfigList.DestinationRules.Items))
	assert.Equal(0, len(istioconfigList.ServiceEntries))
	assert.Equal(0, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeGateways = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(0, len(istioconfigList.VirtualServices.Items))
	assert.Equal(0, len(istioconfigList.DestinationRules.Items))
	assert.Equal(0, len(istioconfigList.ServiceEntries))
	assert.Equal(0, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeVirtualServices = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(0, len(istioconfigList.DestinationRules.Items))
	assert.Equal(0, len(istioconfigList.ServiceEntries))
	assert.Equal(0, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeDestinationRules = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(0, len(istioconfigList.ServiceEntries))
	assert.Equal(0, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeServiceEntries = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(0, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeRules = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(0, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeAdapters = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(1, len(istioconfigList.Adapters))
	assert.Equal(0, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeTemplates = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(1, len(istioconfigList.Adapters))
	assert.Equal(1, len(istioconfigList.Templates))
	assert.Equal(0, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeQuotaSpecs = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(1, len(istioconfigList.Adapters))
	assert.Equal(1, len(istioconfigList.Templates))
	assert.Equal(1, len(istioconfigList.QuotaSpecs))
	assert.Equal(0, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludeQuotaSpecBindings = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(1, len(istioconfigList.Adapters))
	assert.Equal(1, len(istioconfigList.Templates))
	assert.Equal(1, len(istioconfigList.QuotaSpecs))
	assert.Equal(1, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(0, len(istioconfigList.Policies))
	assert.Nil(err)

	criteria.IncludePolicies = true

	istioconfigList, err = configService.GetIstioConfigList(criteria)

	assert.Equal(2, len(istioconfigList.Gateways))
	assert.Equal(2, len(istioconfigList.VirtualServices.Items))
	assert.Equal(2, len(istioconfigList.DestinationRules.Items))
	assert.Equal(1, len(istioconfigList.ServiceEntries))
	assert.Equal(1, len(istioconfigList.Rules))
	assert.Equal(1, len(istioconfigList.Adapters))
	assert.Equal(1, len(istioconfigList.Templates))
	assert.Equal(1, len(istioconfigList.QuotaSpecs))
	assert.Equal(1, len(istioconfigList.QuotaSpecBindings))
	assert.Equal(1, len(istioconfigList.Policies))
	assert.Nil(err)
}

func TestGetIstioConfigDetails(t *testing.T) {
	assert := assert.New(t)

	configService := mockGetIstioConfigDetails()

	istioConfigDetails, err := configService.GetIstioConfigDetails("test", "gateways", "", "gw-1")
	assert.Equal("gw-1", istioConfigDetails.Gateway.Metadata.Name)
	assert.True(istioConfigDetails.Permissions.Update)
	assert.False(istioConfigDetails.Permissions.Delete)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "virtualservices", "", "reviews")
	assert.Equal("reviews", istioConfigDetails.VirtualService.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "destinationrules", "", "reviews-dr")
	assert.Equal("reviews-dr", istioConfigDetails.DestinationRule.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "rules", "", "checkfromcustomer")
	assert.Equal("checkfromcustomer", istioConfigDetails.Rule.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "adapters", "listcheckers", "preferencewhitelist")
	assert.Equal("preferencewhitelist", istioConfigDetails.Adapter.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "templates", "listentries", "preferencesource")
	assert.Equal("preferencesource", istioConfigDetails.Template.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "serviceentries", "", "googleapis")
	assert.Equal("googleapis", istioConfigDetails.ServiceEntry.Metadata.Name)
	assert.Nil(err)

	istioConfigDetails, err = configService.GetIstioConfigDetails("test", "rules-bad", "", "stdio")
	assert.Error(err)
}

func mockGetIstioConfigList() IstioConfigService {
	k8s := new(kubetest.K8SClientMock)
	k8s.On("GetGateways", mock.AnythingOfType("string")).Return(fakeGetGateways(), nil)
	k8s.On("GetVirtualServices", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fakeGetVirtualServices(), nil)
	k8s.On("GetDestinationRules", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fakeGetDestinationRules(), nil)
	k8s.On("GetServiceEntries", mock.AnythingOfType("string")).Return(fakeGetServiceEntries(), nil)
	k8s.On("GetIstioRules", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fakeGetIstioRules(), nil)
	k8s.On("GetAdapters", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fakeGetAdapters(), nil)
	k8s.On("GetTemplates", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fakeGetTemplates(), nil)
	k8s.On("GetQuotaSpecs", mock.AnythingOfType("string")).Return(fakeGetQuotaSpecs(), nil)
	k8s.On("GetQuotaSpecBindings", mock.AnythingOfType("string")).Return(fakeGetQuotaSpecBindings(), nil)
	k8s.On("GetPolicies", mock.AnythingOfType("string")).Return(fakeGetPolicies(), nil)

	return IstioConfigService{k8s: k8s}
}

func fakeGetGateways() []kubernetes.IstioObject {
	gw1 := data.CreateEmptyGateway("gw-1", "test", map[string]string{
		"app": "my-gateway1-controller",
	})

	gw1.GetSpec()["servers"] = []interface{}{
		map[string]interface{}{
			"port": map[string]interface{}{
				"number":   80,
				"name":     "http",
				"protocol": "HTTP",
			},
			"hosts": []interface{}{
				"uk.bookinfo.com",
				"eu.bookinfo.com",
			},
			"tls": map[string]interface{}{
				"httpsRedirect": "true",
			},
		},
	}

	gw2 := data.CreateEmptyGateway("gw-2", "test", map[string]string{
		"app": "my-gateway2-controller",
	})

	gw2.GetSpec()["servers"] = []interface{}{
		map[string]interface{}{
			"port": map[string]interface{}{
				"number":   80,
				"name":     "http",
				"protocol": "HTTP",
			},
			"hosts": []interface{}{
				"uk.bookinfo.com",
				"eu.bookinfo.com",
			},
			"tls": map[string]interface{}{
				"httpsRedirect": "true",
			},
		},
	}

	return []kubernetes.IstioObject{gw1, gw2}
}

func fakeGetVirtualServices() []kubernetes.IstioObject {
	virtualService1 := data.AddRoutesToVirtualService("http", data.CreateRoute("reviews", "v2", 50),
		data.AddRoutesToVirtualService("http", data.CreateRoute("reviews", "v3", 50),
			data.CreateEmptyVirtualService("reviews", "test", []string{"reviews"}),
		),
	)

	virtualService2 := data.AddRoutesToVirtualService("http", data.CreateRoute("details", "v2", 50),
		data.AddRoutesToVirtualService("http", data.CreateRoute("details", "v3", 50),
			data.CreateEmptyVirtualService("details", "test", []string{"details"}),
		),
	)

	return []kubernetes.IstioObject{virtualService1, virtualService2}
}

func fakeGetDestinationRules() []kubernetes.IstioObject {

	destinationRule1 := data.AddSubsetToDestinationRule(data.CreateSubset("v1", "v1"),
		data.AddSubsetToDestinationRule(data.CreateSubset("v2", "v2"),
			data.CreateEmptyDestinationRule("test", "reviews-dr", "reviews")))

	destinationRule1.GetSpec()["trafficPolicy"] = map[string]interface{}{
		"connectionPool": map[string]interface{}{
			"http": map[string]interface{}{
				"maxRequestsPerConnection": 100,
			},
		},
		"outlierDetection": map[string]interface{}{
			"http": map[string]interface{}{
				"consecutiveErrors": 50,
			},
		},
	}

	destinationRule2 := data.AddSubsetToDestinationRule(data.CreateSubset("v1", "v1"),
		data.AddSubsetToDestinationRule(data.CreateSubset("v2", "v2"),
			data.CreateEmptyDestinationRule("test", "details-dr", "details")))

	destinationRule2.GetSpec()["trafficPolicy"] = map[string]interface{}{
		"connectionPool": map[string]interface{}{
			"http": map[string]interface{}{
				"maxRequestsPerConnection": 100,
			},
		},
		"outlierDetection": map[string]interface{}{
			"http": map[string]interface{}{
				"consecutiveErrors": 50,
			},
		},
	}

	return []kubernetes.IstioObject{destinationRule1, destinationRule2}
}

func fakeGetServiceEntries() []kubernetes.IstioObject {
	serviceEntry := kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: "googleapis",
		},
		Spec: map[string]interface{}{
			"hosts": []interface{}{
				"*.googleapis.com",
			},
			"ports": []interface{}{
				map[string]interface{}{
					"number":   443,
					"name":     "https",
					"protocol": "http",
				},
			},
		},
	}
	return []kubernetes.IstioObject{&serviceEntry}
}

func fakeGetIstioRules() []kubernetes.IstioObject {
	stdioRule := kubernetes.GenericIstioObject{}
	stdioRule.Name = "stdio"
	stdioRule.Spec = map[string]interface{}{
		"match": "true",
		"actions": []map[string]interface{}{
			{
				"handler": "handler.stdio",
				"instances": []string{
					"accesslog.logentry",
				},
			},
		},
	}
	return []kubernetes.IstioObject{&stdioRule}
}

func fakeGetAdapters() []kubernetes.IstioObject {
	handler := kubernetes.GenericIstioObject{}
	handler.Name = "preferencewhitelist"
	handler.Spec = map[string]interface{}{
		"overrides": []string{
			"recommendation",
		},
		"blacklist": false,
		"adapter":   "listchecker",
	}
	return []kubernetes.IstioObject{&handler}
}

func fakeGetTemplates() []kubernetes.IstioObject {
	instance := kubernetes.GenericIstioObject{}
	instance.Name = "preferencesource"
	instance.Spec = map[string]interface{}{
		"value":    "source.labels[\"app\"]",
		"template": "listentry",
	}
	return []kubernetes.IstioObject{&instance}
}

func fakeCheckFromCustomerRule() kubernetes.IstioObject {
	checkfromcustomerRule := kubernetes.GenericIstioObject{}
	checkfromcustomerRule.Name = "checkfromcustomer"
	checkfromcustomerRule.Spec = map[string]interface{}{
		"match": "destination.labels[\"app\"] == \"preference\"",
		"actions": []map[string]interface{}{
			{
				"handler": "preferencewhitelist.listchecker",
				"instances": []string{
					"preferencesource.listentry",
				},
			},
		},
	}
	return &checkfromcustomerRule
}

func fakeGetQuotaSpecs() []kubernetes.IstioObject {
	quotaSpec := kubernetes.GenericIstioObject{}
	quotaSpec.Name = "request-count"
	quotaSpec.Spec = map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"quotas": []interface{}{
					map[string]interface{}{
						"charge": 1,
						"quota":  "RequestCount",
					},
				},
			},
		},
	}
	return []kubernetes.IstioObject{&quotaSpec}
}

func fakeGetQuotaSpecBindings() []kubernetes.IstioObject {
	quotaSpec := kubernetes.GenericIstioObject{}
	quotaSpec.Name = "request-count"
	quotaSpec.Spec = map[string]interface{}{
		"quotaSpecs": []interface{}{
			map[string]interface{}{
				"name":      "request-count",
				"namespace": "istio-system",
			},
		},
		"services": []interface{}{
			map[string]interface{}{
				"name": "ratings",
			},
			map[string]interface{}{
				"name": "reviews",
			},
			map[string]interface{}{
				"name": "details",
			},
			map[string]interface{}{
				"name": "productpage",
			},
		},
	}
	return []kubernetes.IstioObject{&quotaSpec}
}

func fakeGetPolicies() []kubernetes.IstioObject {
	policy := kubernetes.GenericIstioObject{}
	policy.Name = "request-count"
	policy.Spec = map[string]interface{}{
		"targets": []interface{}{
			map[string]interface{}{
				"name": "target",
				"port": []interface{}{
					map[string]interface{}{
						"number": 8080,
						"name":   "tomcat",
					},
				},
			},
			map[string]interface{}{
				"name": "target",
				"port": []interface{}{
					map[string]interface{}{
						"number": 80,
						"name":   "nginx",
					},
				},
			},
		},
		"peers": []interface{}{
			map[string]interface{}{
				"mode": "STRICT",
			},
		},
	}
	return []kubernetes.IstioObject{&policy}

}

func fakeGetSelfSubjectAccessReview() []*auth_v1.SelfSubjectAccessReview {
	create := auth_v1.SelfSubjectAccessReview{
		Spec: auth_v1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &auth_v1.ResourceAttributes{
				Namespace: "test",
				Verb:      "create",
				Resource:  "destinationrules",
			},
		},
		Status: auth_v1.SubjectAccessReviewStatus{
			Allowed: true,
			Reason:  "authorized",
		},
	}
	update := auth_v1.SelfSubjectAccessReview{
		Spec: auth_v1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &auth_v1.ResourceAttributes{
				Namespace: "test",
				Verb:      "patch",
				Resource:  "destinationrules",
			},
		},
		Status: auth_v1.SubjectAccessReviewStatus{
			Allowed: true,
			Reason:  "authorized",
		},
	}
	delete := auth_v1.SelfSubjectAccessReview{
		Spec: auth_v1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &auth_v1.ResourceAttributes{
				Namespace: "test",
				Verb:      "delete",
				Resource:  "destinationrules",
			},
		},
		Status: auth_v1.SubjectAccessReviewStatus{
			Allowed: false,
			Reason:  "not authorized",
		},
	}
	return []*auth_v1.SelfSubjectAccessReview{&create, &update, &delete}
}

func mockGetIstioConfigDetails() IstioConfigService {
	k8s := new(kubetest.K8SClientMock)
	k8s.On("GetGateway", "test", "gw-1").Return(fakeGetGateways()[0], nil)
	k8s.On("GetVirtualService", "test", "reviews").Return(fakeGetVirtualServices()[0], nil)
	k8s.On("GetDestinationRule", "test", "reviews-dr").Return(fakeGetDestinationRules()[0], nil)
	k8s.On("GetServiceEntry", "test", "googleapis").Return(fakeGetServiceEntries()[0], nil)
	k8s.On("GetIstioRule", "test", "checkfromcustomer").Return(fakeCheckFromCustomerRule(), nil)
	k8s.On("GetAdapter", "test", "listcheckers", "preferencewhitelist").Return(fakeGetAdapters()[0], nil)
	k8s.On("GetTemplate", "test", "listentries", "preferencesource").Return(fakeGetTemplates()[0], nil)
	k8s.On("GetQuotaSpec", "test", "request-count").Return(fakeGetQuotaSpecs()[0], nil)
	k8s.On("GetQuotaSpecBinding", "test", "request-count").Return(fakeGetQuotaSpecBindings()[0], nil)
	k8s.On("GetSelfSubjectAccessReview", "test", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(fakeGetSelfSubjectAccessReview(), nil)

	return IstioConfigService{k8s: k8s}
}

func TestIsValidHost(t *testing.T) {
	conf := config.NewConfig()
	config.Set(conf)

	virtualServiceIstioObject := kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: "reviews",
		},
		Spec: map[string]interface{}{
			"hosts": []interface{}{
				"reviews",
			},
			"http": []interface{}{
				map[string]interface{}{
					"route": []interface{}{
						map[string]interface{}{
							"destination": map[string]interface{}{
								"host":   "reviews",
								"subset": "v2",
							},
							"weight": 50,
						},
						map[string]interface{}{
							"destination": map[string]interface{}{
								"host":   "reviews",
								"subset": "v3",
							},
							"weight": 50,
						},
					},
				},
			},
		},
	}
	virtualService := models.VirtualService{}
	virtualService.Parse(virtualServiceIstioObject.DeepCopyIstioObject())

	assert.False(t, virtualService.IsValidHost("", ""))
	assert.False(t, virtualService.IsValidHost("", "ratings"))
	assert.True(t, virtualService.IsValidHost("", "reviews"))
}

func TestHasCircuitBreaker(t *testing.T) {
	conf := config.NewConfig()
	config.Set(conf)

	// Note - I don't think the subset definitions here have any impact on the CB
	// detection. They do not do any sort of override so presumably any version, including
	// a v3 would inherit the DR-level CB definition.
	destinationRule1 := kubernetes.GenericIstioObject{
		Spec: map[string]interface{}{
			"host": "reviews",
			"trafficPolicy": map[string]interface{}{
				"connectionPool": map[string]interface{}{
					"http": map[string]interface{}{
						"maxRequestsPerConnection": 100,
					},
				},
				"outlierDetection": map[string]interface{}{
					"http": map[string]interface{}{
						"consecutiveErrors": 50,
					},
				},
			},
			"subsets": []interface{}{
				map[string]interface{}{
					"name": "v1",
					"labels": map[string]interface{}{
						"version": "v1",
					},
				},
				map[string]interface{}{
					"name": "v2",
					"labels": map[string]interface{}{
						"version": "v2",
					},
				},
			},
		},
	}
	dRule1 := models.DestinationRule{}
	dRule1.Parse(destinationRule1.DeepCopyIstioObject())

	assert.False(t, dRule1.HasCircuitBreaker("", "", ""))
	assert.True(t, dRule1.HasCircuitBreaker("", "reviews", ""))
	assert.False(t, dRule1.HasCircuitBreaker("", "reviews-bad", ""))
	assert.True(t, dRule1.HasCircuitBreaker("", "reviews", "v1"))
	assert.True(t, dRule1.HasCircuitBreaker("", "reviews", "v2"))
	assert.True(t, dRule1.HasCircuitBreaker("", "reviews", "v3"))
	assert.False(t, dRule1.HasCircuitBreaker("", "reviews-bad", "v2"))

	destinationRule2 := kubernetes.GenericIstioObject{
		Spec: map[string]interface{}{
			"host": "reviews",
			"subsets": []interface{}{
				map[string]interface{}{
					"name": "v1",
					"labels": map[string]interface{}{
						"version": "v1",
					},
				},
				map[string]interface{}{
					"name": "v2",
					"labels": map[string]interface{}{
						"version": "v2",
					},
					"trafficPolicy": map[string]interface{}{
						"connectionPool": map[string]interface{}{
							"http": map[string]interface{}{
								"maxRequestsPerConnection": 100,
							},
						},
						"outlierDetection": map[string]interface{}{
							"http": map[string]interface{}{
								"consecutiveErrors": 50,
							},
						},
					},
				},
			},
		},
	}
	dRule2 := models.DestinationRule{}
	dRule2.Parse(destinationRule2.DeepCopyIstioObject())

	assert.True(t, dRule2.HasCircuitBreaker("", "reviews", ""))
	assert.False(t, dRule2.HasCircuitBreaker("", "reviews", "v1"))
	assert.True(t, dRule2.HasCircuitBreaker("", "reviews", "v2"))
	assert.False(t, dRule2.HasCircuitBreaker("", "reviews-bad", "v2"))
}

func TestDeleteIstioConfigDetails(t *testing.T) {
	assert := assert.New(t)
	configService := mockDeleteIstioConfigDetails()

	err := configService.DeleteIstioConfigDetail("networking.istio.io", "test", "virtualservices", "", "reviews-to-delete")
	assert.Nil(err)

	err = configService.DeleteIstioConfigDetail("config.istio.io", "test", "templates", "listcheckers", "listchecker-to-delete")
	assert.Nil(err)
}

func mockDeleteIstioConfigDetails() IstioConfigService {
	k8s := new(kubetest.K8SClientMock)
	k8s.On("DeleteIstioObject", "networking.istio.io", "test", "virtualservices", "reviews-to-delete").Return(nil)
	k8s.On("DeleteIstioObject", "config.istio.io", "test", "listcheckers", "listchecker-to-delete").Return(nil)
	return IstioConfigService{k8s: k8s}
}

func TestUpdateIstioConfigDetails(t *testing.T) {
	assert := assert.New(t)
	configService := mockUpdateIstioConfigDetails()

	updatedVirtualService, err := configService.UpdateIstioConfigDetail("networking.istio.io", "test", "virtualservices", "", "reviews-to-update", "{}")
	assert.Equal("test", updatedVirtualService.Namespace.Name)
	assert.Equal("virtualservices", updatedVirtualService.ObjectType)
	assert.Equal("reviews-to-update", updatedVirtualService.VirtualService.Metadata.Name)
	assert.Nil(err)

	updatedTemplate, err := configService.UpdateIstioConfigDetail("config.istio.io", "test", "templates", "listcheckers", "listchecker-to-update", "{}")
	assert.Equal("test", updatedTemplate.Namespace.Name)
	assert.Equal("templates", updatedTemplate.ObjectType)
	assert.Equal("listchecker-to-update", updatedTemplate.Template.Metadata.Name)
	assert.Nil(err)
}

func mockUpdateIstioConfigDetails() IstioConfigService {
	k8s := new(kubetest.K8SClientMock)
	var updatedVirtualService, updatedTemplate kubernetes.IstioObject

	updatedVirtualService = &kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "reviews-to-update",
			Namespace: "test",
		},
	}
	updatedTemplate = &kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "listchecker-to-update",
			Namespace: "test",
		},
	}
	k8s.On("UpdateIstioObject", "networking.istio.io", "test", "virtualservices", "reviews-to-update", mock.AnythingOfType("string")).Return(updatedVirtualService, nil)
	k8s.On("UpdateIstioObject", "config.istio.io", "test", "listcheckers", "listchecker-to-update", mock.AnythingOfType("string")).Return(updatedTemplate, nil)
	return IstioConfigService{k8s: k8s}
}

// mockCreateIstioConfigDetails to verify the behavior of API calls is the same for create and update
func mockCreateIstioConfigDetails() IstioConfigService {
	k8s := new(kubetest.K8SClientMock)
	var createdVirtualService, createdTemplate kubernetes.IstioObject

	createdVirtualService = &kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "reviews-to-update",
			Namespace: "test",
		},
	}
	createdTemplate = &kubernetes.GenericIstioObject{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "listchecker-to-update",
			Namespace: "test",
		},
	}
	k8s.On("CreateIstioObject", "networking.istio.io", "test", "virtualservices", mock.AnythingOfType("string")).Return(createdVirtualService, nil)
	k8s.On("CreateIstioObject", "config.istio.io", "test", "listcheckers", mock.AnythingOfType("string")).Return(createdTemplate, nil)
	return IstioConfigService{k8s: k8s}
}

func TestCreateIstioConfigDetails(t *testing.T) {
	assert := assert.New(t)
	configService := mockCreateIstioConfigDetails()

	createVirtualService, err := configService.CreateIstioConfigDetail("networking.istio.io", "test", "virtualservices", "", []byte("{}"))
	assert.Equal("test", createVirtualService.Namespace.Name)
	assert.Equal("virtualservices", createVirtualService.ObjectType)
	assert.Equal("reviews-to-update", createVirtualService.VirtualService.Metadata.Name)
	assert.Nil(err)

	createTemplate, err := configService.CreateIstioConfigDetail("config.istio.io", "test", "templates", "listcheckers", []byte("{}"))
	assert.Equal("test", createTemplate.Namespace.Name)
	assert.Equal("templates", createTemplate.ObjectType)
	assert.Equal("listchecker-to-update", createTemplate.Template.Metadata.Name)
	assert.Nil(err)
}
