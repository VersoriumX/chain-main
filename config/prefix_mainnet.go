//go:build !testnet
// +build !testnet
{
  "name": "VersoriumX Token List",
  "logoURI": "https://raw.githubusercontent.com/VersoriumX/assets/main/versoriumx-logo.png",
  "keywords": [
    "versoriumx",
    "ethereumx",
    "defi",
    "swap",
    "aggregator",
    "crypto.org",
    "cro"
  ],
  "timestamp": "2024-07-01T12:00:00.000Z",
  "tokens": [
    {
      "chainId": 137,
      "address": "0x8487B97c91ecC1a03b4907B64Bdeab306B888c0E",
      "name": "EthereumX",
      "symbol": "ETHX",
      "decimals": 18,
      "logoURI": "https://raw.githubusercontent.com/VersoriumX/assets/main/ethereumx-logo.png",
      "extensions": {
        "description": "The official utility and governance token for the VersoriumX ecosystem on Polygon."
      }
    },
    {
      "chainId": 394,
      "address": "native",
      "name": "Crypto.org Coin",
      "symbol": "CRO",
      "decimals": 8,
      "logoURI": "https://raw.githubusercontent.com/cosmos/chain-registry/master/cryptoorg/images/cro.svg",
      "extensions": {
        "description": "The native token of the Crypto.org Chain.",
        "coingeckoId": "crypto-com-chain",
        "chainConfig": {
          "coinType": 394,
          "fundraiserPath": "44'/394'/0'/0/1",
          "addressPrefixes": {
            "account": "cro",
            "accountPubkey": "cropub",
            "validator": "crocncl",
            "validatorPubkey": "crocnclpub",
            "consensusNode": "crocnclcons",
            "consensusNodePubkey": "crocnclconspub"
          },
          "coinUnits": {
            "human": "cro",
            "base": "basecro",
            "exponent": 8
          }
        }
      }
    }
  ],
  "version": {
    "major": 1,
    "minor": 1,
    "patch": 0
  }
}
