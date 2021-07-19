package main

import (
	"fmt"
	"strconv"

	"github.com/Philipp15b/go-steamapi"
)

type Steam struct {
	SteamAPIkey string
}

var g Steam

func extInit(args ...string) string {
	if Inited {
		return "Success"
	}
	g = Steam{SteamAPIkey: args[0]}

	Inited = true
	return "Success"
}

func makeUids(args []string) []uint64 {
	var steamIDs []uint64
	for _, elem := range args {
		uid, err := strconv.ParseUint(elem,10,64)
		if err != nil {
			panic(err)
		}
		steamIDs = append(steamIDs, uid)
	}
	return steamIDs
}

func getBans(args ...string) string {
	if !Inited {
		return "Init extension first"
	}
	steamIDs := makeUids(args)
	bans, err := steamapi.GetPlayerBans(steamIDs, g.SteamAPIkey)
	if err != nil {
		return err.Error()
	}
	var res string
	for idx, _ := range bans {
		res += fmt.Sprintf(`["%v",%v,%v,"%s",%v,%d,%d]`, bans[idx].SteamID, bans[idx].CommunityBanned, bans[idx].VACBanned, bans[idx].EconomyBan, bans[idx].NumberOfVACBans, bans[idx].DaysSinceLastBan, bans[idx].NumberOfGameBans)
		if idx != len(bans) - 1 && len(bans) > 0 {
			res += ","
		}
	}
	return fmt.Sprintf("[%s]",res)
}

func getProfileName(args ...string) string {
	if !Inited {
		return "Init extension first"
	}
	steamIDs := makeUids(args)
	names, err := steamapi.GetPlayerSummaries(steamIDs, g.SteamAPIkey)
	if err != nil {
		return err.Error()
	}
	var res string
	for idx, _ := range names {
		res += fmt.Sprintf(`["%v",%d,"%v","%v"]`,
			names[idx].SteamID,
			names[idx].CommunityVisibilityState,
			names[idx].ProfileURL,
			names[idx].PersonaName,
		)
		if idx != len(names) - 1 && len(names) > 0 {
			res += ","
		}
	}
	return fmt.Sprintf("[%s]",res)
}

func getA3Time(args ...string) string {
	return ""
}

func getDateRegister(args ...string) string {
	return ""
}

func getFriends(args ...string) string {
	return ""
}
