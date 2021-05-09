package config

type Discord struct {
	Token   string `config:"discord.token,required"`
	OwnerId string `config:"discord.ownerid"`
}

type Config struct {
	Discord Discord
}
