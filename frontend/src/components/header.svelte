<script lang="ts">
    import { PUBLIC_HOST_URL } from '$env/static/public';

    import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte'
    import { Hamburger } from 'svelte-hamburgers';

    export let isSpotlightOpen: boolean;
    export let isMenuOpen: boolean;

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

<header class="z-[1000] bg-[rgb(26,27,30)] border-b-[0.0625rem] border-dark-500 border-solid h-[70px] max-h-[60px] p-[16px] box-border fixed top-0 left-0 right-0 content-center select-none flex justify-between">  <!-- flex items-center select-none justify-between justify-items-center"> -->
    <a href="/" class="flex items-center hover:scale-[1.01]">
        <img src="{PUBLIC_HOST_URL}/favicon.png" class="h-[35px] w-[35px] inline-block mr-[16px]" alt="skyzar logo"/>
        <p class="text-theme-700 font-bold text-[35px]">
            skyzar
        </p>
    </a>
    <div class="flex space-x-4 items-center max-lg:hidden">
        {#each links as { href, text, external }}
            {#if external}
                <a {href} target="_blank" rel="noopener noreferrer" class="text-text-theme font-semibold text-[20px] hover:scale-[1.02]">
                    {text}
                </a>
            {:else}
                <a {href} class="text-text-theme font-semibold text-[20px] hover:scale-[1.02]">
                    {text}
                </a>
            {/if}
        {/each}
        <button on:click={() => (isSpotlightOpen = !isSpotlightOpen)} class="rounded-3xl px-[16px] border-2 border-theme-700 h-[35px] flex items-center gap-2 hover:border-theme-500 hover:scale-[1.02]">
            <SearchOutline/>
            <p class="text-center">Search</p>
        </button>
    </div>
    <div class="lg:hidden flex items-center">
        <Hamburger bind:open={isMenuOpen} --color="white"/>
    </div>
</header>