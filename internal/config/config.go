package config

import "github.com/sirupsen/logrus"

type Discord struct {
	Token   string `config:"discord.token,required"`
	OwnerId string `config:"discord.ownerid"`
}

type Lavalink struct {
	HttpAddress string `config:"lavalink.httpaddress"`
	WSAddress   string `config:"lavalink.wsaddress"`
	Password    string `config:"lavalink.password"`
}

type Log struct {
	Level logrus.Level `config:"log.level"`
}

type Config struct {
	Debug    bool `config:"debug"`
	Log      Log
	Discord  Discord
	Lavalink Lavalink
}
