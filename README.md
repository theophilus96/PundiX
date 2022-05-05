# PundiX

## API documentation


**GET    /communitypool**

To query all coins in the community pool which is under Governance control

Request
```
GET http://localhost:8080/communitypool
```
Response
```
{
    "pool": [
        {
            "denom": "FX",
            "amount": "35100737564691730392608347.427772988227487229"
        }
    ]
}
```



**GET    /proposal**

Query all governance proposal
Request
```
GET http://localhost:8080/proposal
```
Response
```
{
    "Proposals": [
        {
            "proposal_id": "1",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "Binance Smart Chain support",
                "description": "Summary:\nFunction X foundation is proposing to work on deploying Binance Smart Chain (BSC) onto Function X. This will expand BSC's ecosystem and liquidity into Function X ecosystem (incl PundiXChain, FXCore, etc)\n\nThere are visible benefits to this integration. \n1) [Work in Progress] Once XPOS go on-chain, it will more seamlessly support payments originating from and to BSC. \n2) [Work in Progress] Use cases, tokens and liquidity in Function X ecosystem will be able to move in and out of BSC, including $FX, $PUNDIX, $PURSE and more.\n3) Future chains and DApps on Function X ecosystem will be able to bridge into BSC.\n\nTimeline: \nVoting will last 14 days, if approved the foundation will need  no more than 7 days to deploy. This means a total of 21 days from the day this proposal is submitted for BSC to go live, if voting is in favor FOR this proposal.\n\nVote YES for: BSC will be bridged into Function X. \nVote NO for: will not be bridged into Function X.",
                "params": {
                    "gravity_id": "fx-bsc-bridge",
                    "average_block_time": "5000",
                    "external_batch_timeout": "43200000",
                    "average_external_block_time": "3000",
                    "signed_window": "20000",
                    "slash_fraction": "0.000010000000000000",
                    "oracle_set_update_power_change_percent": "0.001000000000000000",
                    "ibc_transfer_timeout_height": "20000",
                    "oracles": [
                        "fx1wt962ym68xlkxrhfw7z2lrxrc9pws6hglehw33",
                        "fx1rjrwte05yd3tw2005lfgk573nrjmkxcr5yvltx",
                        "fx1757a7a9jdqnlwrysk3crh0jv7dhl85sxt8shtn",
                        "fx1jjpwenetj0u70peyyjy9lewwzjyg2y7suq3hct",
                        "fx1d6xj9yekmpmyafwx9tvgzl2gmzcpe04n54853f"
                    ],
                    "deposit_threshold": {
                        "denom": "FX",
                        "amount": "10000000000000000000000"
                    }
                },
                "chain_name": "bsc"
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "136651785466546095427757830",
                "abstain": "0",
                "no": "0",
                "no_with_veto": "86941150799793632261932"
            },
            "submit_time": "2021-10-08T07:38:18.673436053Z",
            "deposit_end_time": "2021-10-22T07:38:18.673436053Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-10-08T07:41:35.36637132Z",
            "voting_end_time": "2021-10-22T07:41:35.36637132Z"
        },
        {
            "proposal_id": "2",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "Nammy - a cross-browser extension for accessing FX multi-blockchain",
                "description": "Proposal for creation of cross-browser extension for accessing FX multi-blockchain (FX, Ethereum, Bitcoin) in one place.\n\nFeatures:\n- Multi-blockchain access (FX, Ethereum, Bitcoin).\n- Expected support for Chrome, Firefox, Brave and many more.\n- Open source (GPLv3 or later) at GitHub and GitLab.\n- Smart contract integration (e.g. access to variety of Dapps which can be used at hand such as social apps, games, etc.).\n\nThe main development work will span across 12 months (8 months covered by EGF, 4 months self-funded) which will split into 6 milestones (each every 2 months):\n\n1. First release of dependency libraries, command-line tools, initial extension with design and Testnet focused integration.\n2. Release version for Chrome, Firefox (and more) with Testnet focused integration (mainly on FX blockchain). Creation of Nammy node validator.\n3. Release for supported web-browser focusing on FX and ETH blockchains (and more), most likely with initial Dapps integration. Hopefully Mainnet ready (expected over 50 unique users).\n4. First releases (on corresponding web-browser stores) integrated with Mainnet (and Testnet) and Dapps integration. Adding supporting docs. Expected over 200 unique users.\n5. Releases focusing on Dapps with new smart contracts, improved security and easy to use. Documenting usage. Expected over 400 unique installs.\n6. Releases with expanded support for cross-browser compatibility, different blockchains, Dapps and blockchain explorer, and community-driven features. Expected over 500-1000 unique installs.\n\n- Project repository: https://github.com/Nammy-fx\n- Site: https://nammy-fx.github.io/ which is going to be used for releases and downloads (under construction)\n- More details and comments: https://forum.functionx.io/t/a-web-browser-based-fx-wallet-not-a-metamask/1193",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_REJECTED",
            "final_tally_result": {
                "yes": "35654089913691477916477151",
                "abstain": "0",
                "no": "0",
                "no_with_veto": "0"
            },
            "submit_time": "2021-10-11T20:44:17.954377636Z",
            "deposit_end_time": "2021-10-25T20:44:17.954377636Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-10-11T20:44:17.954377636Z",
            "voting_end_time": "2021-10-25T20:44:17.954377636Z"
        },
        {
            "proposal_id": "3",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "f(x)DM - Function X Decentralized Marketing",
                "description": "Summary: this proposal is applying for $1M worth of $FX tokens, which at today's rate 800,000 tokens. If the voting is passed tokens will be drawn from the Ecosystem Genesis Fund into f(x)DM.\n\nFunction X foundation would like to propose a new program called Function X Decentralized Marketing a.k.a f(x)DM. The objective of f(x)DM is to further grow the Function X ecosystem by allowing community members to draw funds for development, marketing and evangelism purposes. Members can draw funds directly from this pot which have (intended) passed governance voting instead of a needing to submit to the Ecosystem Genesis Fund. This allows more effective fund allocations. Funds from f(x)DM marketing needs to be approved by four of seven committee members (four from outside of the foundation). More details: https://forum.functionx.io/t/egf-proposal-request-for-comments-f-x-dm-function-x-decentralized-marketing/1236\n\nTimeline:\nVoting will last 14 days, if approved the foundation will be able to deploy 800,000 FX tokens for 2 years. Vote YES for: f(x)DM will be activated. Vote NO for: f(x)DM will not be activated.",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "119058727536739467390777273",
                "abstain": "0",
                "no": "21016270105100000000000",
                "no_with_veto": "0"
            },
            "submit_time": "2021-10-20T09:41:47.860218179Z",
            "deposit_end_time": "2021-11-03T09:41:47.860218179Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10009000000000000000000"
                }
            ],
            "voting_start_time": "2021-10-20T09:52:00.20045715Z",
            "voting_end_time": "2021-11-03T09:52:00.20045715Z"
        },
        {
            "proposal_id": "4",
            "content": {
                "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
                "title": "f(x)Core parameter change",
                "description": "\n\n\nSummary:\nThis proposal will perform improvements and increase node efficiencies. After the upgrade we will see increased SlashFraction and reduced OracleSetUpdatePowerChangePercentage.\n\nChanges\n1.  Higher slashing rate if malicious nodes are found.\n2. Better on-chain voting power management, without constantly recording voting power changes.\n\nTimeline:\nVoting will last 14 days, if approved the foundation will need no more than 7 days to deploy these changes. This means a total of 21 days from the day this proposal is submitted, if voting is in favor FOR this proposal.\n\nVote YES for: improvement changes.\nVote NO for: no improvement changes.",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "119127823345299929938570481",
                "abstain": "41000000000000000000000",
                "no": "0",
                "no_with_veto": "22609294908370000000000"
            },
            "submit_time": "2021-10-23T06:41:51.315677042Z",
            "deposit_end_time": "2021-11-06T06:41:51.315677042Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-10-23T06:41:51.315677042Z",
            "voting_end_time": "2021-11-06T06:41:51.315677042Z"
        },
        {
            "proposal_id": "5",
            "content": {
                "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
                "title": "Increase FXCore max validator to 50",
                "description": "Summary:\nThis proposal will increase maximum validator count from the original max 20 to 50. It is the first max validator increase since the launch of FXCore and will allow public-hosted validator to join and help secure the network.\n\nChanges\n1. 20 to 50 max validators.\n\nTimeline:\nVoting will last 14 days, if approved max validator will increase to 50.\n\nVote YES for: improvement changes.\nVote NO for: no improvement changes.",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "121132752485299982287384157",
                "abstain": "0",
                "no": "0",
                "no_with_veto": "0"
            },
            "submit_time": "2021-11-11T00:57:15.080917639Z",
            "deposit_end_time": "2021-11-25T00:57:15.080917639Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-11-11T00:57:15.080917639Z",
            "voting_end_time": "2021-11-25T00:57:15.080917639Z"
        },
        {
            "proposal_id": "6",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "TRON support",
                "description": "Summary:\nFunction X foundation is proposing to deploy TRON onto Function X. This will expand TRON's ecosystem and liquidity into Function X ecosystem (incl. PundiXChain, FXCore, etc)\n\nTimeline:\nVoting will last 14 days, if approved the foundation will take 7 days to deploy, hence a total of 21 days from the day this proposal is submitted for TRON to go live, if voting is in favor FOR this proposal.\n\nVote YES for: TRON will be bridged into Function X.\nVote NO for: TRON will not be bridged into Function X.",
                "params": {
                    "gravity_id": "fx-tron-bridge",
                    "average_block_time": "5000",
                    "external_batch_timeout": "43200000",
                    "average_external_block_time": "3000",
                    "signed_window": "20000",
                    "slash_fraction": "0.001000000000000000",
                    "oracle_set_update_power_change_percent": "0.100000000000000000",
                    "ibc_transfer_timeout_height": "20000",
                    "oracles": [
                        "fx1gmhfhmscngtxvdcsm0753txnz5vp8kzjmhchcj",
                        "fx19f7lft94qkkkj0ht4kg5nz3rmrwtzarqfdvaj3",
                        "fx19s08ghjh5g55x2ffqj6c0l378zshcvskxjacra",
                        "fx15m5esmf347eq7mvg9438qx5m8l8f0564y7d8zy",
                        "fx1w65el5ta9v7y6h46qmx4cg0jpu5aq24w6lnvjf",
                        "fx1t6qzk2293ukpunruu9npagczc2hepmurh7drtv",
                        "fx15e8f7nwqlk02jhjlccgvnc093ajumetsl9uj9p",
                        "fx1h86slytsgl6g60jvsg3kjfsrmxysmhskjn5pzg",
                        "fx13m0dtdz3yusf365jcy7y57n8ql80s52w39wjfy",
                        "fx1zxsk6wwxskdlxq69hgs52t4enf2mvd59h7ze7v"
                    ],
                    "deposit_threshold": {
                        "denom": "FX",
                        "amount": "10000000000000000000000"
                    }
                },
                "chain_name": "tron"
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "120960529737155430397650638",
                "abstain": "0",
                "no": "10061771409378375711047",
                "no_with_veto": "0"
            },
            "submit_time": "2021-11-16T03:44:43.693636428Z",
            "deposit_end_time": "2021-11-30T03:44:43.693636428Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-11-16T03:44:43.693636428Z",
            "voting_end_time": "2021-11-30T03:44:43.693636428Z"
        },
        {
            "proposal_id": "7",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "Polygon Support",
                "description": "Function X foundation is proposing to deploy Polygon onto Function X. This will expand Polygon's ecosystem and liquidity into Function X ecosystem (incl. PundiXChain, FXCore, etc)\n\nTimeline:\nVoting will last 14 days, if approved the foundation will take 7 days to deploy Polygon, hence a total of 21 days from the day this proposal is submitted for Polygon to go live, if voting is in favor FOR this proposal.\n\nVote YES for: Polygon will be bridged into Function X.\nVote NO for: Polygon will not be bridged into Function X.",
                "params": {
                    "gravity_id": "fx-polygon-bridge",
                    "average_block_time": "5000",
                    "external_batch_timeout": "43200000",
                    "average_external_block_time": "2000",
                    "signed_window": "20000",
                    "slash_fraction": "0.001000000000000000",
                    "oracle_set_update_power_change_percent": "0.100000000000000000",
                    "ibc_transfer_timeout_height": "20000",
                    "oracles": [
                        "fx1rhzm7e52yu7xnauvk59ymjwhthy0mh7myql5nh",
                        "fx19k0v0zkkzqgjpadezfdxy93wzfaunxcexerqwz",
                        "fx18xrn8stc5j75tf86wutmkangr2g7veyjp8hkzg",
                        "fx1kyarvm4zdelxw0lxyxwrxrnm5j74fwaau8dsd9",
                        "fx1c6npw4s5g2gsnp0gteverjxemyxjrm3ml7k5sp",
                        "fx1ggcz3gp680vn54ezld9vs9p4u7u024mez5d6ga",
                        "fx1lmgre5nekjcphvsh60q0zacqp5c362mhvyw9vs",
                        "fx1fkjyznu3yhuxtqp4mjurpaa79ct4sdxp6sacxk",
                        "fx1de3ysnnwr7wstcy8vgkzc57cat4zdskntaw6ms",
                        "fx1z80r3lle20f8unpc4h99sye6xnwajvj5gsj6j3"
                    ],
                    "deposit_threshold": {
                        "denom": "FX",
                        "amount": "10000000000000000000000"
                    }
                },
                "chain_name": "polygon"
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "121018551912661438509120219",
                "abstain": "0",
                "no": "0",
                "no_with_veto": "0"
            },
            "submit_time": "2021-11-16T03:59:11.061756105Z",
            "deposit_end_time": "2021-11-30T03:59:11.061756105Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2021-11-16T03:59:11.061756105Z",
            "voting_end_time": "2021-11-30T03:59:11.061756105Z"
        },
        {
            "proposal_id": "8",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "Proposal for creation of a NFT collection and following gaming Metaverse on FX(core): ",
                "description": "Proposal for creation of a 3d-model NFT collection and following gaming Metaverse on FX(core): \n\nDAOverse\n\nFeatures: \n-\tGame Ready 3d Models as NFTs of highest Quality\n      o\tPresale with Blindboxproject NFTs (blind boxes act as a voucher)\n      o\tRelease in editions and events every quarter\n-\tCoin DVG that will be used for fueling the metaverse and will be built on FX(core)\n      o\tConnect with FX DEX and other DEX’s, CEX’s\n      o\tConnect with XPOS\n      o\t(one additional governance token may be issued on a later stage if needed)\n-\tMetaverse Profile creation with NFTs for the usage in Social media, Services and Games\n      o\tEvery profile will use FX Wallet address to raise numbers of FXcore users\n      o\tUse of the 16personalities.com test to create unique user social experience\n      o\tMany Services to develop and release (like school (online classes), recreational/leisure activities, concerts, multiplayer games, etc.)\n      o\tOpen to cross chain services at later stages\n-\tMain Game development in Unreal Engine 5\n      o\tBookable storyline written by professional writers\n      o\tPlay +develop +service to earn DVG \n      o\tRPG single player campaign and PVE part\n      o\tE-sports friendly PVP part\n      o\tBuy Land, build houses, meet friends, build a life\n-\tFull Marketing concept\n      o\tEvery marketing option will be used for DAOverse AND FX to raise the user base of FX and the DAOverse\n      o\tBuilding and management of an active Community\n      o\tIncluding planned collaborations with Influencers, NFT-/ Gaming Projects and Gaming Brands \n-\tMarket Maker position of FX\n      o\tEvery FX we have spare and the excess Money we will make will go into FX relevant Market Maker positions\n             \tThis means we only Sell and Buy providing Liquidity and not taking it out of the Markets to stabilize FX \n                    •\tSo if we get the FX for our project you don’t need to have the fear of a price drop. We won’t sell FX by hitting the Market sell button.\n             \tYes this really means we will buy FX back if the Market maker position requests it and our liquid wallet is filled \n                       •\tThis can result in having bought back more FX than actually ever gotten in this proposal\n-\tDAOverse Validator\n      o\tAlready running with full uptime\n      o\t1% regular and 2% Max Commission rate (0% in the starting month)\n\nFX Distribution plan: \nThe Payment 01, 623,000 USD equivalent FX, directly at the end of the Proposal Voting, is mainly needed to build the company and contract the Developers, The Marketing plan will start with finding collaborations / Influencer / Ambassadors and building the community.  Following payments are bound to the progress of the DAOverse and mainly needed for marketing and operations.\nWe propose that every 60 finished NFTs will be a good timing to set up the next payment (50k) for next phase Marketing and development that we can kick off a sellout in July 2022 and from that point on are in our own full monetarization plan and don’t need any more funding from the EGF, since we want to reduce the impact on the EGF to a minimum. \n\n\nFurther investment rounds:\nThe possibility to invest in our work and participate of the earnings will be given by the launch of the DVG Coin which will be released in late 2022 or 2023, or with the participation in our NFT sales.\nWe won’t sell this company to an Investor that will get the power to control us. We already declined two big offers by Game developer companies that wanted to buy our work. \nWe want to build it here as a community focused DAO on FXcore if YOU want us to. \n\nThank you all for your early Support.\n\nYours, \nShady aka PundixSherif/KuzoIV\nPatrick aka 1337Alchimist\n\n\nLinks: \nWhitepaper: \nhttps://daoverse.games/wp-content/uploads/2022/01/Whitepaper-Rev.-0.1-Function-x-EGF-Fund-Proposal-Version.pdf\nDAOverse Validator \nhttps://explorer.functionx.io/fxcore/validator/fxvaloper12dzp2wydkc8rtjeejehsck8kwmrxz7tk3t8h6s\nForum proposal \nhttps://forum.functionx.io/t/daoverse-creation-proposal-welcome-to-the-daoverse/2426\nSocial Media: \nhttps://twitter.com/daoverse_games\nhttps://daoverse.games\nhttps://facebook.com/Daoverse\nhttps://tiktok.com/@daoverse\nhttps://www.youtube.com/channel/UC2RjiuK6uQJDfVLm09NZtBw\nDiscord: https://t.co/KH3Qk5TqL5\n",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_REJECTED",
            "final_tally_result": {
                "yes": "65728590146194226752621993",
                "abstain": "3550052092412264950434553",
                "no": "1290025737593797294481725",
                "no_with_veto": "0"
            },
            "submit_time": "2022-01-10T13:37:38.467362975Z",
            "deposit_end_time": "2022-01-24T13:37:38.467362975Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10031310040000000000000"
                }
            ],
            "voting_start_time": "2022-01-10T13:52:44.230658327Z",
            "voting_end_time": "2022-01-24T13:52:44.230658327Z"
        },
        {
            "proposal_id": "9",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "[Market-Making Initiative from Kronos Research]",
                "description": "Summary: This proposal aims to improve FX liquidity on centralized exchanges. Market makers provide buy/sell orders on either side of the midpoint, allowing users on exchanges to buy, sell, and trade the token more efficiently, giving greater confidence in the asset. In order to provide ample liquidity, Kronos Research is requesting a loan of 6,800,000 FX tokens for a period of 1 year. With this loan, Kronos Research can provide liquidity on Coinbase (FX/USD), Bithumb (FX/KRW), KuCoin (FX/BTC) and Upbit (FX/BTC). Currently the liquidity on these platforms is inconsistent, and we expect, as the terms of this proposal, to increase it by 50%-100% with stable uptimes.\n\nNote: Kronos Research does not attempt to influence market price in either direction.\n\nKronos Background:\nKronos Research is a proprietary trading firm founded in 2018 that trades on average more than $5b USD per day across global platforms. In DeFi, Kronos Research is a leading designated market-maker on platforms like dYdX, Matcha, and ParaSwap. Kronos Research also incubated and is the sole market maker on WOO Network and WOO X.\n\nAction items:\nVote YES for: Accept the proposal\nVote NO for: Reject the proposal",
                "params": {
                    "gravity_id": "",
                    "average_block_time": "",
                    "external_batch_timeout": "",
                    "average_external_block_time": "",
                    "signed_window": "",
                    "slash_fraction": "",
                    "oracle_set_update_power_change_percent": "",
                    "ibc_transfer_timeout_height": "",
                    "oracles": null,
                    "deposit_threshold": {
                        "denom": "",
                        "amount": ""
                    }
                },
                "chain_name": ""
            },
            "status": "PROPOSAL_STATUS_PASSED",
            "final_tally_result": {
                "yes": "144550326512658124745910491",
                "abstain": "12904717058951037480095373",
                "no": "0",
                "no_with_veto": "0"
            },
            "submit_time": "2022-02-12T02:44:27.0832214Z",
            "deposit_end_time": "2022-02-26T02:44:27.0832214Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ],
            "voting_start_time": "2022-02-12T02:44:27.0832214Z",
            "voting_end_time": "2022-02-26T02:44:27.0832214Z"
        }
    ],
    "pagination": {
        "next_key": null,
        "total": "9"
    }
}
```

**GET    /proposal/filtered**

Query all governance proposal
Request but the response is filtered
```
GET http://localhost:8080/proposal/filtered
```
Response
```
{
    "proposals": [
        {
            "proposal_id": "1",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "Binance Smart Chain support"
            },
            "submit_time": "2021-10-08T07:38:18.673436053Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "2",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "Nammy - a cross-browser extension for accessing FX multi-blockchain",
                "recipient": "fx1nammyjr4kcjz75kse8l5lh4ste5u0sh9agh5sh"
            },
            "submit_time": "2021-10-11T20:44:17.954377636Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "3",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "f(x)DM - Function X Decentralized Marketing",
                "recipient": "fx1rjrwte05yd3tw2005lfgk573nrjmkxcr5yvltx"
            },
            "submit_time": "2021-10-20T09:41:47.860218179Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10009000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "4",
            "content": {
                "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
                "title": "f(x)Core parameter change"
            },
            "submit_time": "2021-10-23T06:41:51.315677042Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "5",
            "content": {
                "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
                "title": "Increase FXCore max validator to 50"
            },
            "submit_time": "2021-11-11T00:57:15.080917639Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "6",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "TRON support"
            },
            "submit_time": "2021-11-16T03:44:43.693636428Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "7",
            "content": {
                "@type": "/fx.gravity.crosschain.v1.InitCrossChainParamsProposal",
                "title": "Polygon Support"
            },
            "submit_time": "2021-11-16T03:59:11.061756105Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        },
        {
            "proposal_id": "8",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "Proposal for creation of a NFT collection and following gaming Metaverse on FX(core): ",
                "recipient": "fx1wmqfy7dws99y0e9gg0gwtuu0ympzygjte7654p"
            },
            "submit_time": "2022-01-10T13:37:38.467362975Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10031310040000000000000"
                }
            ]
        },
        {
            "proposal_id": "9",
            "content": {
                "@type": "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
                "title": "[Market-Making Initiative from Kronos Research]",
                "recipient": "fx1e2xjrsa0peksu66dhn5kvma7y8hdyz4zsjlxgn"
            },
            "submit_time": "2022-02-12T02:44:27.0832214Z",
            "total_deposit": [
                {
                    "denom": "FX",
                    "amount": "10000000000000000000000"
                }
            ]
        }
    ]
}
```


**GET    /totalWithdrawnAmount**

To query all coins in the community pool which is under Governance control

Request
```
GET http://localhost:8080/totalWithdrawnAmount
```
Response
```
"90040310040000000000000"
```
 

## FAQ

**Qn: Do we need to git clone the repo before running the Get request?**

A: Nope, you don't have to download the entire repo. You only have to download the Final folder and run Main.go file. That would help start up the HTTP server to which you can send requests.

From the command line in the directory containing main.go, run the code. Use a dot argument to mean “run code in the current directory.”

```
$ go run .
```


**Qn: Should we be replacing the localhost with a known validator IP address?**

A: The validator IP address has already been configured on the backend of the code. Hence, by running the queries such as:
```
GET http://localhost:8080/communitypool
```
We will still be able to retrieve the same results.


## Commonly encountered problems

Do also note that I am using Go 1.18. It is known to cause some errors on Macs. This was caused by an old version of golang.org/x/sys
```
go get -u golang.org/x/sys
```
Running the code above would help solve the issue.