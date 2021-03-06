package install

import (
	"fmt"
	"strings"
)

var (
	GlooNamespacedKinds    []string
	GlooClusterScopedKinds []string
	GlooCrdNames           []string
	GlooECrdNames          []string

	GlooComponentLabels map[string]string
)

func init() {

	GlooComponentLabels = map[string]string{
		"app": "(gloo,glooe-prometheus)",
	}

	GlooNamespacedKinds = []string{
		"Deployment",
		"DaemonSet",
		"Service",
		"ConfigMap",
		"ServiceAccount",
		"Role",
		"RoleBinding",
		"Job",
	}

	GlooClusterScopedKinds = []string{
		"ClusterRole",
		"ClusterRoleBinding",
		"ValidatingWebhookConfiguration",
	}

	GlooCrdNames = []string{
		"gateways.gateway.solo.io",
		"proxies.gloo.solo.io",
		"settings.gloo.solo.io",
		"upstreams.gloo.solo.io",
		"upstreamgroups.gloo.solo.io",
		"virtualservices.gateway.solo.io",
		"routetables.gateway.solo.io",
		"authconfigs.enterprise.gloo.solo.io",
	}

	GlooECrdNames = []string{
		"apidocs.devportal.solo.io",
		"groups.devportal.solo.io",
		"portals.devportal.solo.io",
		"users.devportal.solo.io",
	}

}

func LabelsToFlagString(labelMap map[string]string) (labelString string) {
	for k, v := range labelMap {
		labelString += fmt.Sprintf("%s in %s,", k, v)
	}
	labelString = strings.TrimSuffix(labelString, ",")
	return
}
