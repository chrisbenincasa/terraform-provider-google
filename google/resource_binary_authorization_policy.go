// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func defaultBinaryAuthorizationPolicy(project string) map[string]interface{} {
	return map[string]interface{}{
		"name": fmt.Sprintf("projects/%s/policy", project),
		"admissionWhitelistPatterns": []interface{}{
			map[string]interface{}{
				"namePattern": "gcr.io/google_containers/*",
			},
		},
		"defaultAdmissionRule": map[string]interface{}{
			"evaluationMode":  "ALWAYS_ALLOW",
			"enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
		},
	}
}

func resourceBinaryAuthorizationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceBinaryAuthorizationPolicyCreate,
		Read:   resourceBinaryAuthorizationPolicyRead,
		Update: resourceBinaryAuthorizationPolicyUpdate,
		Delete: resourceBinaryAuthorizationPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBinaryAuthorizationPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_admission_rule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enforcement_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"}, false),
						},
						"evaluation_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"}, false),
						},
						"require_attestations_by": {
							Type:             schema.TypeSet,
							Optional:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: selfLinkNameHash,
						},
					},
				},
			},
			"admission_whitelist_patterns": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_pattern": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"cluster_admission_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enforcement_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY", ""}, false),
						},
						"evaluation_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY", ""}, false),
						},
						"require_attestations_by": {
							Type:             schema.TypeSet,
							Optional:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: selfLinkNameHash,
						},
					},
				},
				Set: func(v interface{}) int {
					// require_attestations_by is a set of strings that can have the format
					// projects/{project}/attestors/{attestor} or {attestor}. We diffsuppress
					// and hash that set on the name, but now we need to make sure that the
					// overall hash here respects that so changing the attestor format doesn't
					// change the hash code of cluster_admission_rules.
					raw := v.(map[string]interface{})

					// modifying raw actually modifies the values passed to the provider.
					// Use a copy to avoid that.
					copy := make((map[string]interface{}))
					for key, value := range raw {
						copy[key] = value
					}
					at := copy["require_attestations_by"].(*schema.Set)
					if at != nil {
						t := convertAndMapStringArr(at.List(), GetResourceNameFromSelfLink)
						copy["require_attestations_by"] = schema.NewSet(selfLinkNameHash, convertStringArrToInterface(t))
					}
					var buf bytes.Buffer
					schema.SerializeResourceForHash(&buf, copy, resourceBinaryAuthorizationPolicy().Schema["cluster_admission_rules"].Elem.(*schema.Resource))
					return hashcode.String(buf.String())
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_policy_evaluation_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ENABLE", "DISABLE", ""}, false),
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBinaryAuthorizationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandBinaryAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	globalPolicyEvaluationModeProp, err := expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(d.Get("global_policy_evaluation_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("global_policy_evaluation_mode"); !isEmptyValue(reflect.ValueOf(globalPolicyEvaluationModeProp)) && (ok || !reflect.DeepEqual(v, globalPolicyEvaluationModeProp)) {
		obj["globalPolicyEvaluationMode"] = globalPolicyEvaluationModeProp
	}
	admissionWhitelistPatternsProp, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(d.Get("admission_whitelist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_whitelist_patterns"); !isEmptyValue(reflect.ValueOf(admissionWhitelistPatternsProp)) && (ok || !reflect.DeepEqual(v, admissionWhitelistPatternsProp)) {
		obj["admissionWhitelistPatterns"] = admissionWhitelistPatternsProp
	}
	clusterAdmissionRulesProp, err := expandBinaryAuthorizationPolicyClusterAdmissionRules(d.Get("cluster_admission_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_admission_rules"); !isEmptyValue(reflect.ValueOf(clusterAdmissionRulesProp)) && (ok || !reflect.DeepEqual(v, clusterAdmissionRulesProp)) {
		obj["clusterAdmissionRules"] = clusterAdmissionRulesProp
	}
	defaultAdmissionRuleProp, err := expandBinaryAuthorizationPolicyDefaultAdmissionRule(d.Get("default_admission_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_admission_rule"); !isEmptyValue(reflect.ValueOf(defaultAdmissionRuleProp)) && (ok || !reflect.DeepEqual(v, defaultAdmissionRuleProp)) {
		obj["defaultAdmissionRule"] = defaultAdmissionRuleProp
	}

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Policy: %#v", obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Policy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceBinaryAuthorizationPolicyRead(d, meta)
}

func resourceBinaryAuthorizationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BinaryAuthorizationPolicy %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	if err := d.Set("description", flattenBinaryAuthorizationPolicyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("global_policy_evaluation_mode", flattenBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(res["globalPolicyEvaluationMode"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("admission_whitelist_patterns", flattenBinaryAuthorizationPolicyAdmissionWhitelistPatterns(res["admissionWhitelistPatterns"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("cluster_admission_rules", flattenBinaryAuthorizationPolicyClusterAdmissionRules(res["clusterAdmissionRules"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("default_admission_rule", flattenBinaryAuthorizationPolicyDefaultAdmissionRule(res["defaultAdmissionRule"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	return nil
}

func resourceBinaryAuthorizationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandBinaryAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	globalPolicyEvaluationModeProp, err := expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(d.Get("global_policy_evaluation_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("global_policy_evaluation_mode"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, globalPolicyEvaluationModeProp)) {
		obj["globalPolicyEvaluationMode"] = globalPolicyEvaluationModeProp
	}
	admissionWhitelistPatternsProp, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(d.Get("admission_whitelist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_whitelist_patterns"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, admissionWhitelistPatternsProp)) {
		obj["admissionWhitelistPatterns"] = admissionWhitelistPatternsProp
	}
	clusterAdmissionRulesProp, err := expandBinaryAuthorizationPolicyClusterAdmissionRules(d.Get("cluster_admission_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_admission_rules"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clusterAdmissionRulesProp)) {
		obj["clusterAdmissionRules"] = clusterAdmissionRulesProp
	}
	defaultAdmissionRuleProp, err := expandBinaryAuthorizationPolicyDefaultAdmissionRule(d.Get("default_admission_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_admission_rule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultAdmissionRuleProp)) {
		obj["defaultAdmissionRule"] = defaultAdmissionRuleProp
	}

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Policy %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
	}

	return resourceBinaryAuthorizationPolicyRead(d, meta)
}

func resourceBinaryAuthorizationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BinaryAuthorizationBasePath}}projects/{{project}}/policy")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	obj = defaultBinaryAuthorizationPolicy(d.Get("project").(string))
	log.Printf("[DEBUG] Deleting Policy %q", d.Id())
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Policy")
	}

	log.Printf("[DEBUG] Finished deleting Policy %q: %#v", d.Id(), res)
	return nil
}

func resourceBinaryAuthorizationPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBinaryAuthorizationPolicyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyAdmissionWhitelistPatterns(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name_pattern": flattenBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(original["namePattern"], d),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRules(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"cluster":                 k,
			"evaluation_mode":         flattenBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(original["evaluationMode"], d),
			"require_attestations_by": flattenBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(original["requireAttestationsBy"], d),
			"enforcement_mode":        flattenBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(original["enforcementMode"], d),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(selfLinkNameHash, v.([]interface{}))
}

func flattenBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["evaluation_mode"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(original["evaluationMode"], d)
	transformed["require_attestations_by"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(original["requireAttestationsBy"], d)
	transformed["enforcement_mode"] =
		flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(original["enforcementMode"], d)
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(selfLinkNameHash, v.([]interface{}))
}

func flattenBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandBinaryAuthorizationPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamePattern, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(original["name_pattern"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamePattern); val.IsValid() && !isEmptyValue(val) {
			transformed["namePattern"] = transformedNamePattern
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRules(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedEvaluationMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(original["evaluation_mode"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["evaluationMode"] = transformedEvaluationMode
		transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(original["require_attestations_by"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
		transformedEnforcementMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(original["enforcement_mode"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["enforcementMode"] = transformedEnforcementMode

		m[original["cluster"].(string)] = transformed
	}
	return m, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = getProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return convertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEvaluationMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(original["evaluation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvaluationMode); val.IsValid() && !isEmptyValue(val) {
		transformed["evaluationMode"] = transformedEvaluationMode
	}

	transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(original["require_attestations_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAttestationsBy); val.IsValid() && !isEmptyValue(val) {
		transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
	}

	transformedEnforcementMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(original["enforcement_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforcementMode); val.IsValid() && !isEmptyValue(val) {
		transformed["enforcementMode"] = transformedEnforcementMode
	}

	return transformed, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = getProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return convertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
