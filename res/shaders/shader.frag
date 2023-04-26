#version 420
#extension GL_ARB_explicit_uniform_location : enable

layout (location = 0) out vec4 frag_colour;
void main() {
    frag_colour = vec4(0.86, 0.76, 0.81, 1.0);
}