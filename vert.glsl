#version 330 core

layout(location = 0) in vec3 vpos;
layout(location = 1) in vec2 texpos;

out vec2 UV;

void main() {
	gl_Position = vec4(vpos, 1);
	UV = texpos;
}
