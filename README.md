## SteamChecker-arma
#### Only for server side using

###### Usage:

- Get your own api key https://steamcommunity.com/dev/apikey
```c
 "steam" callExtension ["init",[API KEY]]
```
- Accepts an unlimited number of steam64 uids, returns the string that you want to convert to an array (parseSimpleArray)
```c
"steam" callExtension ["getBans",["uid1","uid2","uid3"]]
```
	Return: multiple array

		SteamID string
		CommunityBanned bool
		VACBanned bool
		EconomyBan string
		NumberOfVACBans int
		DaysSinceLastBan int
		NumberOfGameBans int

- Accepts an unlimited number of steam64 uids, returns the string that you want to convert to an array (parseSimpleArray)
```c
"steam" callExtension ["getProfileName",["uid1","uid2","uid3"]]
```
	Return: multiple array
		SteamID string
		CommunityVisibilityState int
		ProfileURL string
		PersonaName string

