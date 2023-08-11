<script lang="ts">
	import { Req } from '$lib/wailsjs/go/main/App';
	import { CodeBlock } from '@skeletonlabs/skeleton';
	import ReqInputs from '../components/ReqInputs.svelte';
	import { ConicGradient } from '@skeletonlabs/skeleton';
	import type { ConicStop } from '@skeletonlabs/skeleton';

	let currentMessage = 'https://reqres.in/api/users/2';
	let method = 'GET';
	let body = '';
	let response = '';
	$: console.log(method);
	let pending = false;
	let submitted = false;
	let status: string;

	let bodyProps: string[][] = [[]];
	const conicStops: ConicStop[] = [
		{ color: 'transparent', start: 0, end: 25 },
		{ color: 'rgb(var(--color-primary-500))', start: 75, end: 100 }
	];
	function timeout(ms: number) {
		return new Promise((resolve) => setTimeout(resolve, ms));
	}
	const doRequset = async () => {
		submitted = true;
		pending = true;
		const res = await Req(currentMessage, method, { body: bodyProps });
		console.log(res);
		if (res[0] == '{}') {
			status = res[1];
			pending = false;
			return;
		}
		try {
			const formattedRes = JSON.stringify(JSON.parse(res[0]), null, 4);
			console.log(formattedRes);
			status = res[1];
			response = formattedRes;
			pending = false;
		} catch (error) {
			return;
		} finally {
			pending = false;
		}
	};
</script>

<main class="p-4">
	<form on:submit|preventDefault={doRequset}>
		<div class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-container-token">
			<select bind:value={method} class="text-center">
				<option value="GET" class="text-blue-600">GET</option>
				<option value="POST">POST</option>
				<option value="DELETE">DELETE</option></select
			>
			<input
				bind:value={currentMessage}
				class="bg-transparent border-0 ring-0"
				name="prompt"
				id="prompt"
				placeholder="https://example.com/..."
			/>
			<button class="variant-filled-primary" type="submit">Send</button>
		</div>
	</form>

	<div class="grid grid-cols-2 m-2 gap-2">
		<div class="bg-surface-800 rounded-xl min-h-full">
			<ReqInputs bind:bodyProps />
		</div>

		<div class=" bg-surface-800 rounded-xl min-h-full flex flex-col relative gap-2">
			{#if submitted && pending}
				<div class="w-full absolute -translate-y-1/2 top-1/2">
					<ConicGradient stops={conicStops} spin width="w-16" />
				</div>
			{:else if submitted && !pending}
				<CodeBlock language="json" code={status} class="w-full h-fit my-0 flex-grow-0" />
				<CodeBlock language="json" code={response} class="w-full my-0 " lineNumbers={true} />
			{:else if !submitted}
				<div class="m-auto">Not Sent</div>
			{/if}
		</div>
	</div>
</main>
