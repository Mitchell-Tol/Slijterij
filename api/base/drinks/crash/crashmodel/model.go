package crashmodel

type DrinkCrash struct {
    Ids []string `json:"ids"`
    DropPercentage float32 `json:"drop_percentage"`
}

type CrashResponse struct {
    Updated []string `json:"updated"`
}
