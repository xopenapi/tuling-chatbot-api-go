/*
 * 图灵聊天机器人API
 *
 * 图灵聊天机器人API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tuling
// ChatRsp struct for ChatRsp
type ChatRsp struct {
	Intent Intent `json:"intent,omitempty"`
	Results []ReplyResult `json:"results,omitempty"`
}
