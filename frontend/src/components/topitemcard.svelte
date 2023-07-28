<script lang="ts">
	import type { BazaarTopItem } from "$types";

    export let item: BazaarTopItem;
</script>

<a href={`/product/${item._id}`}>
    {item.display_name}
</a>