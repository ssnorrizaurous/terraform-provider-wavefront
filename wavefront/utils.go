package wavefront

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const (
	accessTypeKey          = "access_type"
	accountsKey            = "account_ids"
	customerKey            = "customer"
	descriptionKey         = "description"
	emailKey               = "email"
	exactMatching          = "EXACT"
	idKey                  = "id"
	lastSuccessfulLoginKey = "last_successful_login"
	nameKey                = "name"
	pageSize               = 100
	permissionsKey         = "permissions"
	prefixesKey            = "prefixes"
	policyRulesKey         = "policy_rules"
	policyTagKey           = "key"
	policyTagValue         = "value"
	roleIdsTagKey          = "role_ids"
	rolesKey               = "roles"
	tagsAndedKey           = "tags_anded"
	tagsKey                = "tags"
	updatedEpochMillisKey  = "updated_epoch_millis"
	updaterIDKey           = "updater_id"
	userGroupsKey          = "user_group_ids"
	userGroupsListKey      = "user_groups"
	usersKey               = "users"
)

// compareStringSliceAnyOrder compares two string slices in any order. It returns
// all the strings appearing only in the left slice followed by all the strings
// appearing only in the right slice.
func compareStringSliceAnyOrder(left, right []string) (onlyLeft, onlyRight []string) {
	leftMap := stringSliceAnyOrderAsMap(left)
	rightMap := stringSliceAnyOrderAsMap(right)
	for str, leftCount := range leftMap {
		rightCount := rightMap[str]
		for leftCount < rightCount {
			onlyRight = append(onlyRight, str)
			rightCount--
		}
		for leftCount > rightCount {
			onlyLeft = append(onlyLeft, str)
			leftCount--
		}
		delete(rightMap, str)
	}
	for str, rightCount := range rightMap {
		for rightCount > 0 {
			onlyRight = append(onlyRight, str)
			rightCount--
		}
	}
	return
}

func stringSliceAnyOrderAsMap(strs []string) map[string]int {
	result := make(map[string]int)
	for _, s := range strs {
		result[s]++
	}
	return result
}

// setStringSlice stores a slice of strings as a set under a particular key.
func setStringSlice(d *schema.ResourceData, key string, strs []string) error {
	result := make([]interface{}, 0, len(strs))
	for _, str := range strs {
		result = append(result, str)
	}
	return d.Set(key, schema.NewSet(schema.HashString, result))
}

// getStringSlice returns a set as a slice of strings.
func getStringSlice(d *schema.ResourceData, key string) []string {
	interfaceList := d.Get(key).(*schema.Set).List()
	result := make([]string, 0, len(interfaceList))
	for _, val := range interfaceList {
		result = append(result, val.(string))
	}
	return result
}

// setStringMap stores a map[string]string under a particular key.
func setStringMap(
	d *schema.ResourceData, key string, strMap map[string]string) error {
	result := make(map[string]interface{}, len(strMap))
	for k, v := range strMap {
		result[k] = v
	}
	return d.Set(key, result)
}

// getStringMap retrieves a map[string]string under a particular key
func getStringMap(d *schema.ResourceData, key string) map[string]string {
	interfaceMap := d.Get(key).(map[string]interface{})
	result := make(map[string]string, len(interfaceMap))
	for k, v := range interfaceMap {
		result[k] = v.(string)
	}
	return result
}

// parseStrArr parses a raw interface from d *schema.ResourceData that contains an array of strings
func parseStrArr(raw interface{}) []string {
	var arr []string
	if raw != nil && len(raw.([]interface{})) > 0 {
		for _, v := range raw.([]interface{}) {
			arr = append(arr, v.(string))
		}
	}
	return arr
}
