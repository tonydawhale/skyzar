<script lang="ts">
    import { PUBLIC_HOST_URL } from '$env/static/public';
    import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte'

    export let isSpotlightOpen: boolean;
</script>

<header class="bg-[rgb(26,27,30)] border-b-[0.0625rem] border-dark-500 border-solid h-[70px] max-h-[60px] p-[16px] box-border fixed top-0 left-0 right-0 inline-grid grid-cols-4 content-center select-none">  <!-- flex items-center select-none justify-between justify-items-center"> -->
    <a href="/" class="flex items-center">
        <img src="{PUBLIC_HOST_URL}/favicon.png" class="h-[35px] w-[35px] inline-block mr-[16px]" alt="skyzar logo"/>
        <p class="text-theme-700 font-bold text-[35px]">
            skyzar
        </p>
    </a>
    <div class="flex space-x-4 items-center justify-center col-span-2">
        <a href="/demand" class="text-text-theme font-semibold text-[20px]">
            Top Demand
        </a>
        <a href="/margin" class="text-text-theme font-semibold text-[20px]">
           Top  Margin ($)
        </a>
        <a href="/margin-percent" class="text-text-theme font-semibold text-[20px]">
            Top Margin (%)
        </a>
        <a href="/crafts" class="text-text-theme font-semibold text-[20px]">
            Top Crafts
        </a>
        <a href="/resell" class="text-text-theme font-semibold text-[20px]">
            Top NPC Resell
        </a>
    </div>
    <div class="flex items-center justify-end">
        <button on:click={() => (isSpotlightOpen = !isSpotlightOpen)} class="rounded-3xl px-[16px] border-2 border-theme-700 h-[35px] flex items-center gap-2">
            <SearchOutline/>
            Search
        </button>
    </div>
</header>