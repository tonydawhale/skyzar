<script lang="ts">
	import type { BazaarTopItem } from "$types";
    import Topitemcard from "$comp/topitemcard.svelte";

    export let data: {items: BazaarTopItem[], count: number, totalProducts: number}
</script>

<div class="mx-[12vw] space-y-[16px] select-none">
    <div class="mx-auto">
        <div class="select-none">
            <p class="text-[50px] font-medium text-center">Top Demand</p>
            <p class="text-[25px] font-light text-center">Viewing <span class="text-theme-700 font-semibold">{data.count}</span> profitable products out of the <span class="text-theme-700 font-semibold">{data.totalProducts}</span> total Hypixel Skyblock Bazaar Items</p>
        </div>
    </div>
    
    <div class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 justify-items-center gap-4 items-stretch">
        {#each data.items as item (item._id)}
            <Topitemcard {item}/>
        {/each}
    </div>
</div>