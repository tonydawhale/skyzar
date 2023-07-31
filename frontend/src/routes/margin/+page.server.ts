import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { API_HOST_URL } from '$env/static/private';
import type { BazaarTopItem } from '$types';

export const load = (async ({ fetch }) => {
    const marginReq = await fetch(`${API_HOST_URL}/bazaar/margin`);

    if (marginReq.status !== 200) {
        return error(marginReq.status, 'Error fetching margin products');
    }

    const data = (await marginReq.json()) as {items: BazaarTopItem[], count: number}

    return data;
}) satisfies PageServerLoad;