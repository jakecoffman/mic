<script>
	const player = document.getElementById('player');

	const handleSuccess = function(stream) {
		if (window.URL) {
			player.srcObject = stream;
		} else {
			player.src = stream;
		}
	};

	function getMedia() {
		navigator.mediaDevices.getUserMedia({ audio: true, video: false })
			.then(handleSuccess)
			.catch(e => console.error('failed to get user media', e))
	}
</script>

<main>
	<h1>Mic permission test</h1>

	<button on:click={getMedia}>Get Media</button>

	<p>Press the button to request media access</p>

	<h2>Recorder/Uploader</h2>
	<input type="file" accept="audio/*" capture>

	<h2>Player</h2>
	<audio id="player" controls></audio>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
