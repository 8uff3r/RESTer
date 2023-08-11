<script lang="ts">
	import { TabGroup, Tab, TabAnchor } from '@skeletonlabs/skeleton';
	let tabSet: number = 0;
	const bodyPropNum = [1];
	export const bodyProps: string[][] = [[]];
</script>

<section>
	<TabGroup>
		<Tab bind:group={tabSet} name="tab1" value={0} class="rounded bg-surface-700 m-4">
			<span>Body</span>
		</Tab>
		<Tab bind:group={tabSet} name="tab2" value={1} class="rounded bg-surface-700 m-4">Params</Tab>
		<Tab bind:group={tabSet} name="tab3" value={2} class="rounded bg-surface-700 m-4">Auth</Tab>
		<!-- Tab Panels --->
		<svelte:fragment slot="panel">
			<div class="p-4 space-y-3">
				{#if tabSet === 0}
					{#each bodyPropNum as num, i}
						<div class="input-group input-group-divider grid-cols-11">
							<input
								type="text"
								placeholder="name"
								class="col-span-5"
								bind:value={bodyProps[i][0]}
							/>
							<input
								type="text"
								placeholder="value"
								class="col-span-5"
								bind:value={bodyProps[i][1]}
							/>
							<button
								class="btn-icon variant-filled place-self-end {num === bodyPropNum.at(-1)
									? '!block'
									: '!hidden'}"
								on:click={() => {
									bodyPropNum.push(num++);
									bodyProps.push([]);
								}}>+</button
							>
						</div>
					{/each}
				{:else if tabSet === 1}
					(tab panel 2 contents)
				{:else if tabSet === 2}
					(tab panel 3 contents)
				{/if}
			</div>
		</svelte:fragment>
	</TabGroup>
</section>
