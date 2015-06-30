package twiliogo

import (
        "encoding/json"
        "net/url"
)

type AvailablePhoneNumberList struct {
        Client                Client
        AvailablePhoneNumbers []IncomingPhoneNumber `json:"available_phone_numbers"`
        URI                   string                `json:"uri"`
}

func GetAvailablePhoneNumberList(client Client, optionals ...Optional) (*AvailablePhoneNumberList, error) {
        var availablePhoneNumberList *AvailablePhoneNumberList

        params := url.Values{}

        for _, optional := range optionals {
                param, value := optional.GetParam()
                params.Set(param, value)
        }

        // U.S. local numbers:
        body, err := client.get(params, client.RootUrl()+"/AvailablePhoneNumbers/US/Local.json")

        if err != nil {
                return nil, err
        }

        availablePhoneNumberList = new(AvailablePhoneNumberList)
        availablePhoneNumberList.Client = client
        err = json.Unmarshal(body, availablePhoneNumberList)

        return availablePhoneNumberList, err
}
