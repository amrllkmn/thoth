<script>
	import { Switch } from '../switch';
	import { Label } from '../label';

	let { checked } = $props();

	let engine = $derived(checked ? 'meilisearch' : 'sqlite');
</script>

<div class="flex items-center space-x-2">
	<Switch id="toggle-meilisearch" bind:checked name="engine" value={engine} />
	<Label for="toggle-meilisearch">Use Meilisearch</Label>
</div>
