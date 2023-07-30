<script lang="ts">
	import type { BazaarTopItem } from "$types";

    export let item: BazaarTopItem;
</script>


<a href={`/product/${item._id}`} class="rounded-lg border-theme-500 border-2 p-[16px] min-w-full hover:scale-[1.02] select-none">
    <div class="flex">
        <h3>{item.display_name}</h3>
    </div>
</a>