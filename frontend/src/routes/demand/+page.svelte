<script lang="ts">
    import {fly} from "svelte/transition";
	import type { BazaarTopItem } from "$types";
    import Topitemcard from "$comp/topitemcard.svelte";

    export let data: {items: BazaarTopItem[], count: number}
</script>

<div class="mx-[12vw]">
    <div class="mx-auto">
        <div class="select-none">
            <p class="text-[50px] font-medium text-center">Top Demand</p>
            <p class="text-[25px] font-light text-center">Viewing <span class="text-theme-700 font-semibold">{data.count}</span> profitable products out of the <span class="text-theme-700 font-semibold">{1253}</span> total Hypixel Skyblock Bazaar Items</p>
        </div>
    </div>
    
    <div class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 justify-items-center gap-4 mt-[16px]">
        {#each data.items as item, i (item._id)}
            <Topitemcard {item}/>
        {/each}
    </div>
</div>