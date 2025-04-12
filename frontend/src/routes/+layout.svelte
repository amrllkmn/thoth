<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<main class="container mx-auto px-4 py-4 sm:py-8">
	<div class="mx-auto max-w-4xl">
		{@render children()}
	</div>
</main>
