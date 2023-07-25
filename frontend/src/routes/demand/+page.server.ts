import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { API_HOST_URL } from '$env/static/private';
import type { BazaarTopItem } from '$types';

export const load = (async ({ fetch }) => {
    const demandReq = await fetch(`${API_HOST_URL}/bazaar/demand`);

    if (demandReq.status !== 200) {
        return error(demandReq.status, 'Error fetching demand products');
    }

    const data = (await demandReq.json()) as {items: BazaarTopItem[], count: number}

    return data;
}) satisfies PageServerLoad;