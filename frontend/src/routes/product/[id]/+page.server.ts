import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { API_HOST_URL } from '$env/static/private';
import type { BazaarItem, BazaarItemHistory } from '$types';

export const load = (async ({ params, fetch }) => {
    const { id } = params;

    const productReq = await fetch(`${API_HOST_URL}/bazaar/products/${id}`);
    const productHistoryReq = await fetch(`${API_HOST_URL}/bazaar/products/${id}/history`);

    if (productReq.status !== 200 || productHistoryReq.status !== 200) {
        return error(productReq.status || productHistoryReq.status, 'Product not found');
    }

    const data = {
        product: (await productReq.json()) as BazaarItem,
        productHistory: (await productHistoryReq.json()) as BazaarItemHistory
    }

    return data;
}) satisfies PageServerLoad;