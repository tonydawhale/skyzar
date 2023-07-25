import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { API_HOST_URL } from '$env/static/private';
import type { BazaarTopItem } from '$types';

export const load = (async ({ fetch }) => {
    const mpReq = await fetch(`${API_HOST_URL}/bazaar/margin_percent`);

    if (mpReq.status !== 200) {
        return error(mpReq.status, 'Error fetching margin percent products');
    }

    const data = (await mpReq.json()) as {items: BazaarTopItem[], count: number}

    return data;
}) satisfies PageServerLoad;