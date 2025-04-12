<script lang="ts">
	import { SearchBar } from '$lib/components/ui/search-bar';
	import { EmptyPage } from '../empty-page';
	import { EngineToggle } from '../engine-toggle';
</script>

<div class="space-y-4 sm:space-y-6">
	<SearchBar />
	<EngineToggle />
	<EmptyPage />
</div>
