<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	let { query } = $props();
</script>

<Input
	type="text"
	name="query"
	bind:value={query}
	placeholder="Search for a book and hit Enter"
	class="placeholder:text-muted-secondary h-10 w-full text-sm sm:text-base"
/>
