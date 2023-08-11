<script lang="ts">
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	// Most of your app wide CSS should be put in this file
	import '../app.postcss';
	import { AppShell, AppBar, Avatar, storeHighlightJs } from '@skeletonlabs/skeleton';
	import hljs from 'highlight.js';
	storeHighlightJs.set(hljs);
</script>

<!-- App Shell -->
<AppShell slotSidebarLeft="bg-black p-16">
	<svelte:fragment slot="sidebarLeft">
		<div class="">Sidebar Left</div>
	</svelte:fragment>
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl">RESTer</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<Avatar src="https://i.pravatar.cc/" width="w-8" class="cursor-pointer" />
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- Page Route Content -->
	<slot />
</AppShell>
