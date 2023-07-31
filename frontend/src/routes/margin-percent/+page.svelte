<script lang="ts">
	import Topitemcard from "$comp/topitemcard.svelte";
	import type { BazaarTopItem } from "$types";

    export let data: {items: BazaarTopItem[], count: number, totalProducts: number}
</script>

<div class="mx-[12vw] space-y-[16px] select-none">
    <div class="mx-auto">
        <div class="select-none">
            <h1 class="text-[50px] font-medium text-center">Top Margin (%)</h1>
            <h2 class="text-[25px] font-light text-center">Viewing <span class="text-theme-700 font-semibold">{data.count}</span> profitable products out of the <span class="text-theme-700 font-semibold">{data.totalProducts}</span> total Hypixel Skyblock Bazaar Items</h2>
        </div>
    </div>
    
    <div class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 justify-items-center gap-4 mt-[16px]">
        {#each data.items as item (item._id)}
            <Topitemcard {item}/>
        {/each}
    </div>
</div>