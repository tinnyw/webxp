import * as THREE from 'three';

const renderScene = () => {
	const canvas = document.querySelector('#mm');
	const renderer = new THREE.WebGLRenderer({ antialias: true, canvas });

	const fov = 90;
	const aspect = 2;  // the canvas default
	const near = 0.1;
	const far = 30;
	const camera = new THREE.PerspectiveCamera(fov, aspect, near, far);
	camera.position.z = 7;

	const scene = new THREE.Scene();
	const boxWidth = 3;
	const boxHeight = 3;
	const boxDepth = 3;
	const geometry = new THREE.BoxGeometry(boxWidth, boxHeight, boxDepth);

	const material = new THREE.MeshPhongMaterial({color: 0x0077CC});  // greenish blue

	const cube = new THREE.Mesh(geometry, material);
	scene.add(cube);

	renderer.render(scene, camera);
	const render = (time) => {
	  time *= 0.0003;  // convert time to seconds
	 
	  cube.rotation.x = time;
	  cube.rotation.y = time;
	 
	  renderer.render(scene, camera);
	 
	  requestAnimationFrame(render);
	}
	requestAnimationFrame(render);

	const color = 0xFFFFFF;
	const intensity = 3;
	const light = new THREE.DirectionalLight(color, intensity);
	light.position.set(-1, 2, 4);
	scene.add(light);
}

renderScene()
