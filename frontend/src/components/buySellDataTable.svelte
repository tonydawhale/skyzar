<script lang="ts">
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
    import type { BazaarItemBuySellSummary } from '$types';

    export let data: BazaarItemBuySellSummary[];
</script>

<Table striped={true} shadow hoverable={true}>
    <TableHead theadClass="bg-theme-700">
        <TableHeadCell>#</TableHeadCell>
        <TableHeadCell># of Orders</TableHeadCell>
        <TableHeadCell># of Items</TableHeadCell>
        <TableHeadCell>Cost Per Unit</TableHeadCell>
        <TableHeadCell>Total Price</TableHeadCell>
    </TableHead>
    <TableBody>
        {#each data as item, idx}
            <TableBodyRow>
                <TableBodyCell>{idx + 1}</TableBodyCell>
                <TableBodyCell>{item.orders.toLocaleString()}</TableBodyCell>
                <TableBodyCell>{item.amount.toLocaleString()}</TableBodyCell>
                <TableBodyCell>{item.pricePerUnit.toLocaleString()} coins</TableBodyCell>
                <TableBodyCell>{(item.pricePerUnit * item.amount * item.orders).toLocaleString()} coins</TableBodyCell>
            </TableBodyRow>
        {/each}
        <TableBodyRow>            
            <TableBodyCell>Total</TableBodyCell>
            <TableBodyCell></TableBodyCell>
            <TableBodyCell>{data.reduce((acc, cur) => acc + cur.amount, 0).toLocaleString()}</TableBodyCell>
            <TableBodyCell></TableBodyCell>
            <TableBodyCell>{data.reduce((acc, cur) => acc + (cur.pricePerUnit * cur.amount * cur.orders), 0).toLocaleString()} coins</TableBodyCell>
        </TableBodyRow>
    </TableBody>
</Table>