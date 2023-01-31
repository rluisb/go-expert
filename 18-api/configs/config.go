package configs

import (
    "github.com/spf13/viper"
    "github.com/go-chi/jwtauth"
)

type conf struct {
    DBDriver string `mapstructure:"DB_DRIVER"`
    DBHost string `mapstructure:"DB_HOST"`
    DBPort string `mapstructure:"DB_PORT"`
    DBUser string `mapstructure:"DB_USER"`
    DBPassword string `mapstructure:"DB_PASSWORD"`
    DBName string `mapstructure:"DB_NAME"`
    WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
    JWTSecret string `mapstructure:"JWT_SECRET"`
    JWTExpiresIn int `mapstructure:"JWT_EXPIRESIN"`
    TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
    var cfg *conf
    viper.SetConfigName("app_config")
    viper.SetConfigType("env")
    viper.AddConfigPath(path)
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    err = viper.Unmarshal(&cfg)
    if err != nil {
        panic(err)
    }
    cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
    return cfg, nil
}