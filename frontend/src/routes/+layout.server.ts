import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { API_HOST_URL } from '$env/static/private';

export const load = (async ({ fetch }) => {
    const products = await fetch(`${API_HOST_URL}/bazaar/products`);

    if (products.status !== 200) {
        return error(products.status, 'Error fetching products');
    }

    return (await products.json()) as { [key: string]: string };
}) satisfies LayoutServerLoad;
