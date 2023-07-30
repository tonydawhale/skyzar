<script lang="ts">
	import type { BazaarItem, BazaarItemHistory } from "$types";
    import BuySellDataTable from "$comp/buySellDataTable.svelte";
    import History24hChart from "$comp/charts/history_24h.svelte";

    export let data: { product: BazaarItem, productHistory: BazaarItemHistory}
</script>

<div class="grid grid-cols-2 gap-[40px] mx-[12vw] select-none">
    <div class="col-span-2">
        <p class="text-center font-semibold text-theme text-[28px] mb-[16px]">{data.product.display_name}</p>
        <div class="border-2 border-theme-700 rounded-lg p-[16px]">
            <div class="text-center font-semibold text-theme text-[20px]">Quick Stats</div>
            <div class="grid sm:grid-cols-1 md:grid-cols-2 gap-[16px]">
                <div>
                    <div class="border-b-2 border-theme-700 text-theme font-semibold text-[18px] text-center my-[16px]">
                        Buy Order Info
                    </div>
                    <div class="px-[12px]">
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Instasell Price</p>
                            <p>{data.product.buy_price.toLocaleString()} coin{data.product.buy_price != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Total Buy Orders</p>
                            <p>{data.product.buy_orders.toLocaleString()} order{data.product.buy_orders != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Total Demand</p>
                            <p>{data.product.buy_volume.toLocaleString()} item{data.product.buy_volume != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Buy Volume</p>
                            <p>{data.product.buy_moving_week.toLocaleString()} item{data.product.buy_moving_week != 1 ? 's': ''}</p>
                        </div>
                    </div>
                </div>
                <div>
                    <div class="border-b-2 border-theme-700 text-theme font-semibold text-[18px] text-center my-[16px]">
                        Sell Order Info
                    </div>
                    <div class="px-[12px]">
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Instabuy Price</p>
                            <p>{data.product.sell_price.toLocaleString()} coin{data.product.sell_price != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Total Sell Orders</p>
                            <p>{data.product.sell_orders.toLocaleString()} order{data.product.sell_orders != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Total Supply</p>
                            <p>{data.product.sell_volume.toLocaleString()} item{data.product.sell_volume != 1 ? 's': ''}</p>
                        </div>
                        <div class="flex justify-between text-theme font-light text-[16px]">
                            <p>Sell Volume</p>
                            <p>{data.product.sell_moving_week.toLocaleString()} item{data.product.sell_moving_week != 1 ? 's': ''}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="max-lg:col-span-2">
        <p class="font-bold text-[28px] text-theme text-center">Buy Orders</p>
        <BuySellDataTable data={data.product.buy_data} />
    </div>
    <div class="max-lg:col-span-2">
        <p class="font-bold text-[28px] text-theme text-center">Sell Orders</p>
        <BuySellDataTable data={data.product.sell_data} />
    </div>
    <div class="col-span-2 rounded-lg border-2 border-theme-700 p-[16px]">
        <p class="font-bold text-[28px] text-theme text-center">24H Price History</p>
        <History24hChart data={data.productHistory.history_24h} />
    </div>
</div>