<script lang="ts">
	import type { BazaarTopItem } from "$types";

    export let item: BazaarTopItem;
</script>

<a href={`/product/${item.hypixel_product_id}`}>
    {item._id}
</a>