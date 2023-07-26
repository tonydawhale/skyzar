<script lang="ts">
	import type { BazaarTopItem } from "$types";
    import Topitemcard from "$comp/topitemcard.svelte";

    export let data: {items: BazaarTopItem[], count: number}
</script>

<div class="select-none">
    <p class="text-[50px] font-medium text-center text-text-theme">Top Demand</p>
    <p class="text-[25px] font-light text-center text-text-theme">Viewing <span class="text-theme-700 font-semibold">{data.count}</span> profitable products out of the <span class="text-theme-700 font-semibold">{1253}</span> total Hypixel Skyblock Bazaar Items</p>
</div>
<div class="grid gric-cols-3">
    {#each data.items as item}
        <Topitemcard {item}/>
    {/each}
</div>