package types

type InstalledPlugins struct {
	Traefik bool
}

type PluginMetaData struct {
	AppMetaData
	Name string
}
