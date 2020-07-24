package bnb

//todo: need to update param before deploying
const (
	// fixed params
	DenomBNB              = "BNB"
	MinConfirmationsBlock = 3
	MaxOrphanBlocks       = 1000

	// mainnet
	MainnetBNBChainID         = "Binance-Chain-Tigris"
	MainnetTotalVotingPowers  = 11000000000000
	MainnetURLRemote          = "https://seed1.longevito.io:443"
	MainnetGenesisBlockHeight = 79394120

	// local
	//TestnetBNBChainID         = "Binance-Dev"
	//TestnetTotalVotingPowers  = 1000000000000
	//TestnetURLRemote          = "http://localhost:26657"
	//TestnetGenesisBlockHeight = 1000
	//TestnetGenesisHeaderStr   = "eyJoZWFkZXIiOnsidmVyc2lvbiI6eyJibG9jayI6MTAsImFwcCI6MH0sImNoYWluX2lkIjoiQmluYW5jZS1EZXYiLCJoZWlnaHQiOjEwMDAsInRpbWUiOiIyMDIwLTAzLTI4VDEyOjUwOjI3LjEwMDU5M1oiLCJudW1fdHhzIjowLCJ0b3RhbF90eHMiOjEsImxhc3RfYmxvY2tfaWQiOnsiaGFzaCI6IjRBMzFFMDU3MUM5N0M1NkE2OTgwRDQ1OTlENEFCNjY4MDVCMjI0ODYwNjBDQkMyRTA0MkRFNjg5RkJBODRCMUMiLCJwYXJ0cyI6eyJ0b3RhbCI6MSwiaGFzaCI6IjJFQTlCMTdEMzI1MDVFQjU2QTEwQjcwOUFDNDVFRDQyQjk0QjAwM0QxRTRBMzFCOTAwMzE5OEVEMDM1MDM1MDIifX0sImxhc3RfY29tbWl0X2hhc2giOiJGNDVGMDkxNTE2NjM4NDlGMjlBMURFM0FCMkRGNjM2NkQzOTEzNjU0QjQxQjAxRDVDNTZGNTcwRDgzMEMyNkU0IiwiZGF0YV9oYXNoIjoiIiwidmFsaWRhdG9yc19oYXNoIjoiRTcxQzcxNEJGOEI4RTYyOUE0MjY2RTY0RTJCQUU1QURBMTUxODVCRUU1QTI2MTcxRENCQzc2NUFDRDQ0RDZGMyIsIm5leHRfdmFsaWRhdG9yc19oYXNoIjoiRTcxQzcxNEJGOEI4RTYyOUE0MjY2RTY0RTJCQUU1QURBMTUxODVCRUU1QTI2MTcxRENCQzc2NUFDRDQ0RDZGMyIsImNvbnNlbnN1c19oYXNoIjoiMjk0RDhGQkQwQjk0Qjc2N0E3RUJBOTg0MEYyOTlBMzU4NkRBN0ZFNkI1REVBRDNCN0VFQ0JBMTkzQzQwMEY5MyIsImFwcF9oYXNoIjoiNkYwNTZDOTA2RkFGRjE2NDAxNzQ3OUMyQTY3OEYyNkY0MzQxQkNEOTFCRDcxNEVEQThDNkZBNDJGMzhCNEM0NiIsImxhc3RfcmVzdWx0c19oYXNoIjoiIiwiZXZpZGVuY2VfaGFzaCI6IiIsInByb3Bvc2VyX2FkZHJlc3MiOiI4N0U3MzM0MjI5NjY2ODVDMUIyNEY0MkEzMTg0QUM5NTlFQzQ5QTRDIn0sImRhdGEiOnsidHhzIjpudWxsfSwiZXZpZGVuY2UiOnsiZXZpZGVuY2UiOm51bGx9LCJsYXN0X2NvbW1pdCI6eyJibG9ja19pZCI6eyJoYXNoIjoiNEEzMUUwNTcxQzk3QzU2QTY5ODBENDU5OUQ0QUI2NjgwNUIyMjQ4NjA2MENCQzJFMDQyREU2ODlGQkE4NEIxQyIsInBhcnRzIjp7InRvdGFsIjoxLCJoYXNoIjoiMkVBOUIxN0QzMjUwNUVCNTZBMTBCNzA5QUM0NUVENDJCOTRCMDAzRDFFNEEzMUI5MDAzMTk4RUQwMzUwMzUwMiJ9fSwicHJlY29tbWl0cyI6W3sidHlwZSI6MiwiaGVpZ2h0Ijo5OTksInJvdW5kIjowLCJibG9ja19pZCI6eyJoYXNoIjoiNEEzMUUwNTcxQzk3QzU2QTY5ODBENDU5OUQ0QUI2NjgwNUIyMjQ4NjA2MENCQzJFMDQyREU2ODlGQkE4NEIxQyIsInBhcnRzIjp7InRvdGFsIjoxLCJoYXNoIjoiMkVBOUIxN0QzMjUwNUVCNTZBMTBCNzA5QUM0NUVENDJCOTRCMDAzRDFFNEEzMUI5MDAzMTk4RUQwMzUwMzUwMiJ9fSwidGltZXN0YW1wIjoiMjAyMC0wMy0yOFQxMjo1MDoyNy4xMDA1OTNaIiwidmFsaWRhdG9yX2FkZHJlc3MiOiI4N0U3MzM0MjI5NjY2ODVDMUIyNEY0MkEzMTg0QUM5NTlFQzQ5QTRDIiwidmFsaWRhdG9yX2luZGV4IjowLCJzaWduYXR1cmUiOiJkRERSUWlrcUdERHBkK3A4NDQwTFdDRUlpNVdqWHhwTmZ1WStaTVZ0d0NPeGlodHlGVEdlVjFIS3lSRCtsUUZCVlVBekkyU1NUKzNURXdDdHRwb0FDdz09In1dfX0="

	// testnet
	TestnetBNBChainID         = "Binance-Chain-Ganges"
	TestnetTotalVotingPowers  = 11000000000000
	TestnetURLRemote          = "https://data-seed-pre-0-s3.binance.org:443"
	TestnetGenesisBlockHeight = 79473100
)
