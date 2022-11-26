package util

type UpdateValues struct {
	Name        string
	Plural      string
	ExampleName string

	Resource string

	Project    bool
	Namespace  string
	YAMLObject string
}

const templateUpdateCurl = `
### Update {{ .Name }}

First retrieve the current object into a file ` + "`object.yaml`" + `. This could look like:
` + "```yaml" + `
{{ .YAMLObject }}
` + "```" + `

Run the following curl command to update a single ` + myObjectInNamespace + `:
` + "```bash" + `
# Replace the {{ .ExampleName }} in the url below with the name of the {{ .Name }} you want to update
curl -s "https://$LOFT_DOMAIN/kubernetes/management/apis/management.loft.sh/v1/` + curlNamespace + `{{ .Resource }}/{{ .ExampleName }}" \
     -X PUT --insecure \
     -H "Content-Type: application/yaml" \
     -H "Authorization: Bearer $ACCESS_KEY" \
     --data-binary "$(cat object.yaml)"
` + "```" + `

### Patch {{ .Name }}

Patching a resource is useful if you want to generically exchange only a small portion of the object instead of retrieving the whole object first and then modifying it.
To learn more about patches in Kubernetes, please take a look at the [official docs](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/#use-a-json-merge-patch-to-update-a-deployment).

Run the following curl command to add a new annotation ` + "`my-annotation: my-value`" + ` to the ` + myObjectInNamespace + ` via a patch:
` + "```bash" + `
# Replace the {{ .ExampleName }} in the url below with the name of the {{ .Name }} you want to update
curl -s "https://$LOFT_DOMAIN/kubernetes/management/apis/management.loft.sh/v1/` + curlNamespace + `{{ .Resource }}/{{ .ExampleName }}" \
     -X PATCH --insecure \
     -H "Content-Type: application/json-patch+json" \
     -H "Authorization: Bearer $ACCESS_KEY" \
     --data '[{"op": "add", "path": "/metadata/annotations/my-annotation", "value": "my-value"}]'
` + "```" + `
`

const templateUpdateKubectl = `
### Update {{ .Name }}

Run the following command to update ` + myObjectInNamespace + `:
` + "```bash" + `
kubectl edit {{ .Resource }}.management.loft.sh {{ .ExampleName }} ` + kubectlNamespace + `
` + "```" + `

Then edit the object and upon save, kubectl will update the resource.

### Patch {{ .Name }}

Patching a resource is useful if you want to generically exchange only a small portion of the object instead of retrieving the whole object first and then modifying it.
To learn more about patches in Kubernetes, please take a look at the [official docs](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/#use-a-json-merge-patch-to-update-a-deployment).

Run the following kubectl command to add a new annotation ` + "`my-annotation: my-value`" + ` to the ` + myObjectInNamespace + ` via a patch:
` + "```bash" + `
kubectl patch {{ .Resource }}.management.loft.sh {{ .ExampleName }} ` + kubectlNamespace + ` \
        --type json \
        -p '[{"op": "add", "path": "/metadata/annotations/my-annotation", "value": "my-value"}]'
` + "```" + `
`

const TemplateUpdateObject = `
import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

You can either use curl or kubectl to update {{ .Plural }}.

<Tabs
    defaultValue="kubectl"
    values={[
      {label: 'kubectl', value: 'kubectl'},
      {label: 'curl', value: 'curl'},
    ]}>
  <TabItem value="kubectl">

` + templateUpdateKubectl + `

  </TabItem>
  <TabItem value="curl">

` + templateUpdateCurl + `

  </TabItem>
</Tabs>
`
