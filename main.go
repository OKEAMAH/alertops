package main

import (
	"context"
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
	config "github.com/aliceh/alertops/pkg/config"
	pd "github.com/aliceh/alertops/pkg/pagerduty"
	utils "github.com/aliceh/alertops/pkg/utils"
)

func main() {

	config, err := config.LoadConfig(config.Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	c, err := pd.NewConfig(config.Token, config.Teams, config.SilentUser, config.IgnoredUsers)
	if err != nil {
		fmt.Println(err)
		return
	}
	Users := utils.DifferenceOfSlices(c.TeamsMemberIDs, config.IgnoredUsers)
	currentUser, _ := c.Client.GetUserWithContext(ctx, c.CurrentUser.ID, pagerduty.GetUserOptions{})
	fmt.Printf("%v\n", currentUser.Name)

	fmt.Printf("Testing GetAlerts command\n")
	highAcknowledgedIncidents, _ := pd.HighAcknowledgedIncidents(c.Client, Users)
	for _, inc := range highAcknowledgedIncidents.Incidents {
		id := inc.ID
		alerts, _ := pd.GetAlerts(c.Client, id, pagerduty.ListIncidentAlertsOptions{})
		fmt.Printf("\n,%v ,\n", alerts)
	}
	fmt.Printf("\n")
	fmt.Printf("END Testing GetAlerts command\n")

	a, err := pd.GetIncidents(c.Client, pagerduty.ListIncidentsOptions{})
	fmt.Printf("GetIncidents Output:  %v\n %v\n", a, err)
	// triggered_incidents, err := c.Client.GetCurrentUserWithContext(ctx, pagerduty.GetCurrentUserOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("%+v", triggered_incidents)

}
