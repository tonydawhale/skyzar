<script lang="ts">
	import Topitemcard from "$comp/topitemcard.svelte";
	import type { BazaarTopItem } from "$types";

    export let data: {items: BazaarTopItem[], count: number}
</script>

<div class="mx-auto">
    <div class="select-none">
        <p class="text-[50px] font-medium text-center">Top Margin (%)</p>
        <p class="text-[25px] font-light text-center">Viewing <span class="text-theme-700 font-semibold">{data.count}</span> profitable products out of the <span class="text-theme-700 font-semibold">{1253}</span> total Hypixel Skyblock Bazaar Items</p>
    </div>
</div>

<div class="grid gric-cols-3">
    {#each data.items as item}
        <Topitemcard {item}/>
    {/each}
</div>