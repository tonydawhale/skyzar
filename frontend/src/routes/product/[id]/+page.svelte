<script lang="ts">
	import type { BazaarItem, BazaarItemHistory } from "$types";

    export let data: { product: BazaarItem, productHistory: BazaarItemHistory}

    console.log(data)
</script>

<div class="justify-items-center items-center">
    {data.product.hypixel_product_id}
</div>