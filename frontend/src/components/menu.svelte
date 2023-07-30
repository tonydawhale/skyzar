<script lang="ts">
    import { fly, scale } from 'svelte/transition';
    import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte'

    export let open: boolean;
    export let isSpotlightOpen: boolean;

    let links = [
        {
            href: '/demand',
            text: 'Demand'
        },
        {
            href: '/margin',
            text: 'Margin ($)'
        },
        {
            href: '/margin-percent',
            text: 'Margin (%)'
        },
        {
            href: 'https://github.com/tonydawhale/skyzar',
            text: 'Github',
            external: true
        },
        {
            href: 'https://ko-fi.com/tonydawhale',
            text: 'Donate',
            external: true
        }
    ]
</script>

{#if open}
    <div class="flex flex-col items-center z-[999] min-h-[100vh] min-w-[100vw] pt-[86px] fixed bg-[rgb(26,27,30)]" transition:fly={{y: -15, delay: 50}}>
        {#each links as {href, text, external} (text)}
            {#if external}
                <a {href} target="_blank" rel="noopener noreferrer" class="text-text-theme font-semibold text-[20px] hover:scale-[1.02] my-[16px] text-center" on:click={() => (open = !open)}>
                    {text}
                </a>
            {:else}
                <a {href} class="text-text-theme font-semibold text-[20px] hover:scale-[1.02] my-[16px] text-center" on:click={() => (open = !open)}>
                    {text}
                </a>
            {/if}
        {/each}
        <button on:click={() => (isSpotlightOpen = !isSpotlightOpen)} class="rounded-3xl px-[16px] border-2 border-theme-700 h-[35px] flex items-center gap-2 hover:border-theme-500 hover:scale-[1.02]">
            <SearchOutline/>
            <p class="text-center">Search</p>
        </button>
    </div>
{/if}