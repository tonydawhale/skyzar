<script lang="ts">
	import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte'
    import SvelteSpotlight from 'svelte-spotlight/SvelteSpotlight.svelte';
    import { matchSorter } from 'match-sorter';

    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';

    export let isOpen: boolean;
	export let isMenuOpen: boolean;
    let query = "";

    let items = [
        {
            title: 'REVENANT_CATALYST'
        }
    ];
    $: results = matchSorter(items, query, { keys: ['title'] });
</script>

<SvelteSpotlight
	{results}
	bind:query
	combo={browser && !/mac/i.test(navigator.platform)
		? { key: 'j', ctrlKey: true }
		: { key: 'j', metaKey: true }}
	bind:isOpen
	searchPlaceholder="Search for a product..."
	modalClass={'w-[600px] max-w-[95%] bg-dark-600 shadow-lg overflow-hidden rounded-md'}
	headerClass={'relative leading-[1.55rem] box-border block'}
	inputClass={'pl-[3.125rem] rounded-b-[0.0625rem] border-solid border-[rgb(44,46,51)] h-[3.125rem] appearance-none w-full block text-left min-h-[3.125rem] rounded-t-[0.25rem] bg-dark-600 focus:outline-none'}
	resultIdKey="title"
	on:select={(event) => {
        goto(`/product/${event.detail.title}`)
		isOpen = false
		isMenuOpen = false
	}}
>
	<div slot="headerLeft" class="items-center justify-center absolute top-0 bottom-0 left-0 flex z-[1] w-[3.125rem] border-0">
		<SearchOutline class="h-[16px] w-[16px]"/>
	</div>
	<div
		slot="result"
		let:selected
		let:result
		class={`hover:bg-dark-200  cursor-pointer text-sm px-10 py-3 w-full ${
			selected ? 'bg-dark-300' : ''
		} `}
	>
		{result.title}
		<!-- <p class="text-slate-500 text-sm">{result.description}</p> -->
	</div>

	<div slot="noResults" class="px-10 py-3">
		<p class="text-slate-500 text-sm">No results...</p>
	</div>
</SvelteSpotlight>