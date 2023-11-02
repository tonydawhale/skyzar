<script lang="ts">
	import type { BazaarTopItem } from "$types";

    export let item: BazaarTopItem;
    let baseItemImageUrl = `https://sb.tonydoes.dev/assets`

    let itemImageUrl: string;

    if (item._id.toLowerCase().startsWith("enchantment_")) itemImageUrl = `${baseItemImageUrl}/item/ENCHANTED_BOOK`
    else if (item._id.toLowerCase().startsWith("essence_")) itemImageUrl = `${baseItemImageUrl}/essence/${item._id.toLowerCase().replace("essence_", "")}`
    else itemImageUrl = `${baseItemImageUrl}/item/${item._id}`
</script>


<a href={`/product/${item._id}`} class="rounded-lg border-theme-500 border-2 p-[16px] min-w-full hover:scale-[1.02] select-none">
    <div class="flex flex-row items-center gap-[20px] mb-4">
        <img src={itemImageUrl} class="w-[50px] h-[50px] rounded-lg" alt={item.display_name}/>
        <h3 class="font-bold text-xl">{item.display_name}</h3>
    </div>
    <div class="grid sm:grid-cols-1 md:grid-cols-2 justify-center gap-[20px]">
        <div>
            <h4>Sell Price: {item.buy_price.toLocaleString()} coins</h4>
            <h4>Buy Volume: {item.buy_volume.toLocaleString()} items</h4>
        </div>
        <div>
            <h4>Buy Price: {item.sell_price.toLocaleString()} coins</h4>
            <h4>Sell Volume: {item.sell_volume.toLocaleString()} items</h4>
        </div>
        <div class="col-span-2">
            <h4 class="text-center">Margin: {item.margin.toLocaleString()} coins ({(item.margin_percent * 100).toLocaleString()}%)</h4>
        </div>
    </div>
</a>