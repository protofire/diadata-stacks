const { ApiPromise, WsProvider } = require("@polkadot/api");
const ethers = require("ethers");
const bignumber = ethers.BigNumber;

async function collateralCurrencies(api) {
  const collateralCurrencies =
    await api.query.vaultRegistry.systemCollateralCeiling.entries();
  const collateralCurrenciesmap = new Map();
  collateralCurrencies.map((collateralCurrencie) => {
    let key = collateralCurrencie[0].toHuman();
    let value = collateralCurrencie[1].toHuman();
    // TODO filter only ibtc as wrapped {"collateral":{"Token":"DOT"},"wrapped":{"Token":"IBTC"}}
    // console.log("key",key[0].collateral.Token);
    // console.log("value",JSON.stringify(value));

    collateralCurrenciesmap.set(key[0].collateral.Token, value);
  });
  return collateralCurrenciesmap;
}

//https://github.com/DefiLlama/DefiLlama-Adapters/blob/f0bcc4e65ebe43ce00c77951376a2249f2931dc3/projects/interlay-collateral/api.js

async function totalUserVaultCollateral(api) {
  const totalUserVaultCollaterals =
    await api.query.vaultRegistry.totalUserVaultCollateral.entries();
  const totalUserVaultCollateralsmap = new Map();

  totalUserVaultCollaterals.map((totalUserVaultCollateral) => {
    let key = totalUserVaultCollateral[0].toHuman();
    let value = totalUserVaultCollateral[1].toHuman();

    totalUserVaultCollateralsmap.set(key[0].collateral.Token, value);
  });
  return totalUserVaultCollateralsmap;
}

async function oracleaggregrate(api) {
  const oracleaggregrators = await api.query.oracle.aggregate.entries();

  const oracleaggregratemap = new Map();

  oracleaggregrators.map((oracleaggregrate) => {
    let key = oracleaggregrate[0].toHuman();
    let value = oracleaggregrate[1].toHuman();

    if (key[0].ExchangeRate) {
      console.log("key,", key);
      console.log("value,", value);
      oracleaggregratemap.set(key[0].ExchangeRate.Token, value);
    }
  });
  return oracleaggregratemap;
}

async function totalIssuance(api) {
  const totalIssuances = await api.query.tokens.totalIssuance.entries();
  const totalIssuancesemap = new Map();

  totalIssuances.map((totalIssuance) => {
    let key = totalIssuance[0].toHuman();

    let value = totalIssuance[1].toHuman();

    console.log("key", key);
    console.log("value", value);
    totalIssuancesemap.set(key[0].Token, value);
  });
  return totalIssuancesemap;
}

async function setup() {
  const wsProvider = new WsProvider(
    "wss://interlay.api.onfinality.io/public-ws"
  );
  const api = await ApiPromise.create({
    provider: wsProvider,
  });

  let collateralCurrenciesmap = new Map();
  let totalUserVaultCollateralmap = new Map();
  let oracleaggregatormap = new Map();
  let totalIssuancesemap = new Map();

  collateralCurrenciesmap = await collateralCurrencies(api);
  console.log("collateralCurrencies", collateralCurrenciesmap);

  totalUserVaultCollateralmap = await totalUserVaultCollateral(api);
  console.log("totalUserVaultCollateral", totalUserVaultCollateralmap);

  oracleaggregatormap = await oracleaggregrate(api);
  console.log("oracleaggregrate", oracleaggregatormap);

  totalIssuancesemap = await totalIssuance(api);
  console.log("totalIssuance", totalIssuancesemap);

  /*

total_backable = collateral_currencies.map(|currency| vaultRegistry.totalUserVaultCollateral[currency] / oracle.aggregate[currency]).sum();
total_issued = tokens.totalIssuance[IBTC];
assert!(total_issued < total_backable);

*/

  let total_backable = bignumber.from(0);

  for (let [collateralCurrency, value] of collateralCurrenciesmap) {
    console.log("collateralCurrency", collateralCurrency);

    let totalUserVaultCollateralcurrency = bignumber.from(
      totalUserVaultCollateralmap.get(collateralCurrency).replaceAll(",", "")
    );
    let oracleaggregatecurrency = bignumber.from(
      oracleaggregatormap.get(collateralCurrency).replaceAll(",", "")
    );
    console.log(
      "totalUserVaultCollateralcurrenct",
      totalUserVaultCollateralcurrency.toString()
    );

    oracleaggregatecurrency = oracleaggregatecurrency.div(1e12);
    oracleaggregatecurrency = oracleaggregatecurrency.div(1e6); //TODO while doing 1e18 its crashing somehow.

    let currentcurrencydecimal = 10;
    let btcdecimal = 8;
    // TODO  oracleaggregatecurrency/currentcurrencydecimal ie  (10^10 / 10^8)
    oracleaggregatecurrency = oracleaggregatecurrency.div(1e2);

    totalUserVaultCollateralcurrency =
      totalUserVaultCollateralcurrency.div(1e10);

    total_backable = total_backable.add(
      totalUserVaultCollateralcurrency.div(oracleaggregatecurrency)
    );
  }

  let total_issued = bignumber.from(
    totalIssuancesemap.get("IBTC").replaceAll(",", "")
  );
  total_issued = total_issued.div(1e8);

  console.log("total_backable", total_backable.toString());
  console.log("total_issued", total_issued.toString());

  await api.disconnect();
}

setup();
