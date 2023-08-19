<script lang="ts">
	import type { BazaarItemHistoryData } from '$types';
    import { Line } from 'svelte-chartjs'

    import {
        Chart as ChartJS,
        Title,
        Tooltip,
        Legend,
        LineElement,
        LinearScale,
        PointElement,
        CategoryScale,
    } from 'chart.js';

    ChartJS.register(
        Title,
        Tooltip,
        Legend,
        LineElement,
        LinearScale,
        PointElement,
        CategoryScale
    ); 

    export let data: BazaarItemHistoryData[];
    export let labelType: 'time' | 'date' = 'time'
    export let limit: number = 0;

    $: history_data = limit > 0 ? data.slice(limit > data.length ? 0 : data.length - limit, data.length) : data

    $: chartData = {
        labels: labelType === 'time' ? history_data.map((item) => {
            return new Date(item.time * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
        }) : history_data.map((item) => {
            return new Date(item.time * 1000 - 1000).toLocaleDateString([], { month: 'short', day: 'numeric' })
        }),
        datasets: [
            {
                label: 'Buy Price',
                data: history_data.map((item) => {
                    return item.buy_price
                }),
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
                pointRadius: 0,
                tension: 0.1
            },
            {
                label: 'Sell Price',
                data: history_data.map((item) => {
                    return item.sell_price
                }),
                borderColor: 'rgb(53, 162, 235)',
                backgroundColor: 'rgba(53, 162, 235, 0.5)',
                pointRadius: 0,
                tension: 0.1
            }
        ]
    }
</script>

<Line
    data={chartData}
    options={{
        plugins: {
            filler: {
                propagate: true
            },
            legend: {
                display: true,
                position: "top"
            },
        },
        interaction: {
            mode: "index",
            axis: "xy",
            intersect: false
        },
        datasets: {
            line: {
                pointRadius: 0
            }
        },
        elements: {
            point: {
                radius: 0
            }
        }
    }}
/>