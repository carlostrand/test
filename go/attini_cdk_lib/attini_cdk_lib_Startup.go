// Attini resources
package attini_cdk_lib


type Startup struct {
	Commands *map[string]interface{} `field:"optional" json:"commands" yaml:"commands"`
	CommandsTimeout *float64 `field:"optional" json:"commandsTimeout" yaml:"commandsTimeout"`
}

