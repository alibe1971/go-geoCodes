package geoCodes

type Config struct {
    Settings Settings `json:"settings"`
}

type Settings struct {
    Languages ConfigLanguages `json:"languages"`
}

type ConfigLanguages struct {
    InPackage map[string]string `json:"inPackage"`
    Default   string    `json:"default"`
}


