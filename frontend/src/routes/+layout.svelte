<script lang="ts">
    import '../app.css';
    import { page } from '$app/stores';
    import { PUBLIC_HOST_URL } from '$env/static/public';

    import Header from '$comp/header.svelte';
	import Spotlight from '$comp/spotlight.svelte';

    let isSpotlightOpen = false;

    import "nprogress/nprogress.css";
    import NProgress from "nprogress";
    import { navigating } from "$app/stores";

    NProgress.configure({
        minimum: 0.16,
    });

    $: {
        if ($navigating) {
            NProgress.start();
        } else NProgress.done();
    }
</script>

<svelte:head>
    <title>Skyzar: The #1 Free & Open-Source Bazaar Interface</title>
    <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
    <link rel="icon" type="image/svg+xml" href="{PUBLIC_HOST_URL}/favicon.svg"/>
    <link rel="icon" type="image/png" href="{PUBLIC_HOST_URL}/favicon.png"/>
    <meta name="theme-color" content="#19AAD5"/>
    <meta property="og:site_name" content="Skyzar.app"/>
    <meta property="og:description" content="Skyzar provides up to date Hypixel Skyblock Bazaar product prices, historical price data, and other information all for free."/>
    <meta property="og:title" content="skyzar.app"/>
	<meta name="author" content="Skyzar Team" />
	<meta name="robots" content="index, follow" />
	<meta property="og:url" content={$page.url.toString()} />

</svelte:head>

<div class="block box-border">
    <Header bind:isSpotlightOpen />
    <div class="flex box-border">
        <main class="grow shrink w-[100vw] box-border min-h-[100vh] pt-[86px] px-[16px] pb-[16px]">
            <Spotlight bind:isOpen={isSpotlightOpen} />
            <slot />
        </main>
    </div>
</div>